#include <filesystem>
#include <vector>
#include <string>
#include <iostream>
#include <fstream>

/// @class List all posible language convertors.
class Convert{
	public:
		/// Add new language to the list.
		/// @param ext Is the extension.
		/// @param lang Is the language in markdown.
		inline void addLanguage(std::string ext, std::string lang){
			extensions_.push_back(std::pair(ext,lang));
		}
		/// Get the whole list.
		/// @return Whole list.
		inline const std::vector<std::pair<std::string, std::string>>& extensions(){
			return extensions_;
		}
		inline void printLanguages(){
			std::cout << " Supported extensions:" << std::endl;
			for(auto&& [ext, lang] : extensions_){
				std::cout << "  * " << lang << " with extension " << ext << std::endl;
			}
		}
	private:
		/// List of languages.
		std::vector<std::pair<std::string, std::string>> extensions_;
};

/// Print help.
void printHelp(){
	std::cout << "Simple converter between project source files and markdown file." << std::endl;
	std::cout << " Usage: codemd input [flags ...] [optional export file	]" << std::endl;
	std::cout << "  -r --reverse : to create a project from markdown file." << std::endl;
	std::cout << "  -h --help : see this help." << std::endl;
	std::cout << "  Otherwise input is a root directory of a project." << std::endl;
}

/// Create the database.
Convert createConvert(){
	Convert C;
	C.addLanguage(".py","python");
	C.addLanguage(".cpp","cpp");
	C.addLanguage(".java","java");
	C.addLanguage(".cd", "cs");
	C.addLanguage(".c", "c");
	C.addLanguage(".h", "c");
	C.addLanguage(".hpp","cpp");
	C.addLanguage(".hs","hs");
	C.addLanguage(".pl","pl");
	C.addLanguage(".sh","sh");
	return C;
}

/// Print contents of file to given stream.
void printFile(const std::filesystem::path& p, const std::string& lang, std::ostream& os){
	std::ifstream ifs;
	ifs.open(p);
	std::string line;
	os << "## File: `" << p << "`" << std::endl << std::endl;
	os << "```" << lang << std::endl;
	bool first = true;
	while(getline(ifs,line, ' ')){
		if(first) os << line;
		else os << ' ' << line;
		first = false;
	}
	os << "```" << std::endl << std::endl;
	
}

/// Go through all directories and find files.
void goThrough(const std::string& filePath, const std::string& ext, const std::string& lang, std::ostream& os){
	namespace fs = std::filesystem;
	fs::path root (filePath);
	for(fs::recursive_directory_iterator it (root); it != fs::end(it) ; ++it){
		fs::path file = (*it);
		if(file.extension() == ext){
			printFile(file, lang, os);
		}
	}
}

/// Reverse the process.
void reverse(const std::string& filePath){
	namespace fs = std::filesystem;
	std::ifstream ifs;
	fs::path file(filePath);
	fs::path root = file.parent_path();
	ifs.open(file);
	std::string line;
	std::ofstream ofs;
	bool write = false;
	while(getline(ifs,line)){
		if(line == "```"){
			write = false;
			ofs.close();
		}
		else if(write){
			ofs << line << std::endl;
		}
		else if(line.substr(0,11) == "## File: `\""){
			fs::path tmpfile = line.substr(11,line.size() - 13);
			fs::path whole = root / tmpfile;
			fs::create_directories(whole.parent_path());
			ofs.open(whole);
		}
		else if(line.substr(0,3) == "```"){
			write = true;
		}			
	}
}

/// Read arguments.
std::pair<bool, bool> readArgs(std::string& ofile, const std::vector<std::string>& args){
	bool first = true;
	bool re = false;
	bool help = false;
	for(auto&& arg : args){
		if(arg == "-r" || arg == "--reverse")
			re = true;
		else if(arg == "-h" || arg == "--help")
			help = true;
		else if(!first)
			ofile = arg;
		first = false;
	}
	return {re, help};
}

int main(int argc, char** argv){
	try{
		Convert C = createConvert();
		std::vector<std::string> args (1+argv,argc+argv);
		if(args.size() == 0) return 1;
		std::string ofile = "export.md";
		auto [re, help] = readArgs(ofile, args);
		if(help){
			printHelp();
			C.printLanguages();
			return 0;
		}
		if(re){
			reverse(args[0]);
			return 0;
		}
		std::ofstream ofs;
		ofs.open(ofile);
		ofs << "";
		for(auto&& pair : C.extensions())
			goThrough(args[0], pair.first, pair.second, ofs);
		return 0;
	} catch (std::exception& e){
		std::cerr << e.what() << std::endl;
	} catch(...){
		std::cerr << "Unknown exception caught." << std::endl;
	}
}
