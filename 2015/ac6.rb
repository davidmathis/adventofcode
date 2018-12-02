input = File.read('i6.txt')

@lights = Hash.new

(0..999).each do |col|
    (0..999).each do |row|
	@lights["#{col},#{row}"] = false
    end
end

input.split("\n").each do |line|
    case line.split(" ")[0] 
	when "turn"
	    action = line.split(" ")[1]
	    start_col,start_row = line.split(" ")[2].split(",")[0],line.split(" ")[2].split(",")[1]
            finish_col,finish_row = line.split(" ")[4].split(",")[0],line.split(" ")[4].split(",")[1]
	when "toggle"
	    action = line.split(" ")[0] 
	    start_col,start_row = line.split(" ")[1].split(",")[0],line.split(" ")[1].split(",")[1]
	    finish_col,finish_row = line.split(" ")[3].split(",")[0],line.split(" ")[3].split(",")[1]
    end
    (start_col.to_i..finish_col.to_i).each do |col|
	(start_row.to_i..finish_row.to_i).each do |row|
	    if action == "toggle"
		@lights["#{col},#{row}"] = !@lights["#{col},#{row}"]
	    elsif action == "on"
		@lights["#{col},#{row}"] = true
	    elsif action == "off"
		@lights["#{col},#{row}"] = false
	    end
	end
    end
end

puts @lights.values.count(true)
