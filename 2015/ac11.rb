input = "cqjxjnds"
bad_c = ['i','o','l']
my_pw = input

def has_alpha(word)
    alpha = ('a'..'z').to_a
    (0..word.size-3).each do |i|
	if alpha[alpha.index(word[i]) + 1] == word[i+1]
	    if alpha[alpha.index(word[i]) + 2] == word[i + 2]
		return true
	    end
        end
    end
    return false
end	

until my_pw.scan(/(.)\1/).uniq.size > 1 && !bad_c.any? {|x| my_pw.include?(x)} && has_alpha(my_pw)
	my_pw.succ!
end

p "Current: #{my_pw}"

my_pw.succ!

until my_pw.scan(/(.)\1/).uniq.size > 1 && !bad_c.any? {|x| my_pw.include?(x)} && has_alpha(my_pw)
	        my_pw.succ!
end

puts "Next: #{my_pw}"
