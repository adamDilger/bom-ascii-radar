#!/bin/bash

if [[ ! -a /tmp/bom-ascii-background.png ]]; then
 curl 'http://www.bom.gov.au/products/radar_transparencies/IDR763.background.png' \
 	-H 'Referer: http://www.bom.gov.au/products/IDR763.loop.shtml' \
 	-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36' \
 	--insecure \
 	-o /tmp/bom-ascii-background.png;
fi

echo "FUK"

bg_color="gray"

magick /tmp/bom-ascii-background.png \
	-fill "$bg_color" \
	-opaque "#C08000" \
	-opaque "#E0D8B8" \
	-fill "transparent" \
	-opaque "#C0D8E8" \
	background.png;
