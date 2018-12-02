input = File.readlines('i7.txt').map(&:chomp)
actions = { 'AND' => '&', 'OR' => '|', 'NOT' => '~', 'LSHIFT' => '<<', 'RSHIFT' => '>>' }
done = false

signals = input.map do |line|
	      line.gsub!(/([a-z]+)/, "@\\1")
	      signal, gate = line.split(' -> ')
    	      actions.each {|k,v| signal.sub!(k,v)}
	      if gate == "@b"
	          "#{gate} = 16076"
	      else
	          "#{gate} = #{signal}"
	      end
	  end

until done
    done = true
    signals.each {|line| eval(line) rescue done = false}
end

puts @a 
