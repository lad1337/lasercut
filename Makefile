KICAD_SYMBOL_DIR=/Applications/KiCad/KiCad.app/Contents/SharedSupport/symbols

.PHONY: clean
clean:
	rm -f kle2netlist.erc kle2netlist.log kle2netlist_lib_sklib.py lasercut-converted.json

.PHONY: convert
convert:
	# npm install @ijprest/kle-serial --save
	./pcb/kle-convert.js lasercut.json


.PHONY: netlist
netlist: convert
	# https://github.com/lad1337/kle2netlist
	KICAD_SYMBOL_DIR=/Applications/KiCad/KiCad.app/Contents/SharedSupport/symbols kle2netlist --layout lasercut-converted.json --switch-library perigoso/keyswitch-kicad-library --no-xml --lib-path ./pcb/ --lib-path ./pcb/keebio-components --output-dir pcb --name lastercut --switch-footprint Hotswap --controller-circuit promicro

	
.PHONY: traces
traces:
	#-us global -is prioritized
	/usr/local/opt/openjdk@11/bin/java -jar ~/bin/freerouting-1.5.0.jar -de ~/workspace/lasercut/pcb/lasercut.dsn -us global -is prioritized
