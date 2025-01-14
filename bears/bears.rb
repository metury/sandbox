$dices = ["\u{2680}", "\u{2681}", "\u{2682}", "\u{2683}", "\u{2684}", "\u{2685}"]
$bears = [0, 0, 2, 0, 4, 0]

def get_dices()
  ret = [0, 0, 0, 0]
  (0..3).each do |i|
	tmp = rand(0..5)
	ret[i] = [$dices[tmp], $bears[tmp]]
  end
  return ret
end

$guesed = 0

loop do
  values = get_dices()
  puts "How many bears are near the lake?\n#{values[0][0]} #{values[1][0]}\n#{values[2][0]} #{values[3][0]}"
  sum = 0
  values.each do |val|
	sum += val[1]
  end
  input = gets.chomp
  break if input.downcase == "exit"
  if sum == Integer(input)
	puts "That is correct answer."
        $guesed += 1
  else
	puts "No, correct answer is #{sum}."
        $guesed = 0
  end
  if $guesed >= 6
    puts "You probably know the trick. You already guesed #{$guesed} times in a row."
  end
end

puts "Bears went to sleep.."
