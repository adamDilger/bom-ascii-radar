#!/bin/bash

# base="http://www.bom.gov.au/products/radar_transparencies/IDR764.topography.png"
# background="http://www.bom.gov.au/products/radar_transparencies/IDR764.topography.png"
#
# curl 'http://www.bom.gov.au/products/radar_transparencies/IDR764.topography.png' \
#   -H 'Accept: image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8' \
#   -H 'Accept-Language: en-GB,en-US;q=0.9,en;q=0.8' \
#   -H 'Cache-Control: no-cache' \
#   -H 'Connection: keep-alive' \
#   -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36' \
#   --insecure -o "background.png"


convert IDR083.background.png -set colorspace RGB -set colorspace Gray computed_background.png

tput cup 0 0

ascii-image-converter computed_background.png -b --threshold 200

sleep 0.2

tput cup 0 0

while true; do
	for filename in ./IDR08A*.png; do
		composite -compose dst-atop computed_background.png $filename - \
			| ascii-image-converter - -b -C

		sleep 1
		tput cup 0 0
	done
done


# while true; do
# 	ascii-image-converter IDR764.T.202403040829.png -b
# 	sleep 1
# 	clear
#
# 	ascii-image-converter IDR764.T.202403040834.png -b
# 	sleep 1
# 	clear
#
# 	ascii-image-converter IDR764.T.202403040839.png -b
# 	sleep 1
# 	clear
#
# 	ascii-image-converter IDR764.T.202403040844.png -b
# 	sleep 1
# 	clear
#
# 	ascii-image-converter IDR764.T.202403040849.png -b
# 	sleep 1
# 	clear
# done;
