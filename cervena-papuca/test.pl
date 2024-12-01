#!/usr/bin/perl

our $papuc = "cervena-papuca.go";
our $res_file = "RESULTS";
our @tests = (15, 20, 30, 40, 50, 60, 70, 80, 90);
our $repetitions = 50;
our $gp = "plot.gp";

open(my $in, '>', $res_file) or die $!;
print $in "";
close($in);

open(my $plot, '>', $gp) or die $!;
print $plot "set title \"Červená papuč - hra\"\n";
print $plot "set xlabel \"Počet karet\"\n";
print $plot "set ylabel \"Délka hry\"\n";
print $plot "set terminal png\n";
print $plot "set output 'plot.png'\n";
print $plot "plot \"RESULTS\" using 1:2 title \"Náhodná generace balíku.\"\n";
close($plot);

foreach my $test (@tests) {
	for my $i (1..$repetitions) {
		open(my $in, '>>', $res_file) or die $!;
		print $in "$test ";
		close($in);
		system("go run $papuc $test | tail -n 1 | awk '{print \$4}' >> $res_file")
	}
}

system("gnuplot $gp");
