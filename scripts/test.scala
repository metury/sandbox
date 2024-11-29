//> using scala 3.5.2

@main
def hello(): Unit =
	println("Ahoj!")
	println("Tohle je SCALA!!!")
	val lines = scala.io.Source.fromFile("test.scala").getLines
	for (line <- lines) println(line)
