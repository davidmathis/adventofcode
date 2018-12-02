input = File.read('i8.txt')

total = 0

input.each_line() do |line| 
    line.chomp!
    total += line.inspect.length - line.length
end

puts "Total: #{total}"
