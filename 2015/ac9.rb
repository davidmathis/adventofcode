input = File.readlines('i9.txt').map(&:chomp)
trips = {}
my_trips = {}

input.each do |line|
    s,f,d = line.split(' ').values_at(* line.split(' ').each_index.select {|i| i.even?})
    trips[[s,f].sort] = d.to_i
end

trips.keys.sort.flatten.uniq.permutation.map do |trip|
    my_trip = 0
    (0..trip.size-2).each do |i|
	routes = trips.select {|k,v| k.include? trip[i] and k.include? trip[i+1] }
	my_trip += routes.values[0]
    end
    my_trips[trip] = my_trip
end

p "Best trip is #{my_trips.sort_by(&:last)[0]}"
p "Worst trip is #{my_trips.sort_by(&:last)[-1]}"
