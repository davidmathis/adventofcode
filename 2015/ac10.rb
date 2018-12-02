input = "1113222113"
output = ""
i, n, x = 0, 1, 1

until x > 50
    while i < input.size
    	if input[i] == input[i+1]
	    n += 1
	else
   	    output << n.to_s + input[i]
	    n = 1
        end
        i += 1
    end
    i = 0
    input = output
    output = ""
    x +=1 
end

puts input.size
