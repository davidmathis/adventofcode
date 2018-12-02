require 'digest/md5'
input = "yzbqklnj"
i = 0

while true
    break if Digest::MD5.hexdigest(input + i.to_s)[0..4] == "00000"
    i += 1
end

puts i
