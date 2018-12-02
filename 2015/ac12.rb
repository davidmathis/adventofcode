input = File.read('i12.txt')

puts input.scan(/-?\d+/).map(&:to_i).reduce(:+) 
