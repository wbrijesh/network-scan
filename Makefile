PROJECT_NAME := network-scan
PLATFORMS := darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64 windows/386 windows/amd64
EXECUTABLES_DIR := executables

.PHONY: all build setup clean

all: setup build

build:
	@mkdir -p $(EXECUTABLES_DIR)
	@for platform in $(PLATFORMS); do \
		IFS='/' read -r -a array <<< "$$platform"; \
		GOOS=$${array[0]}; \
		GOARCH=$${array[1]}; \
		if [ $$GOOS = "windows" ]; then \
			OUTPUT_NAME=$(PROJECT_NAME)_$${GOOS}_$${GOARCH}.exe; \
		else \
			OUTPUT_NAME=$(PROJECT_NAME)_$${GOOS}_$${GOARCH}; \
		fi; \
		echo "Building for $$GOOS $$GOARCH"; \
		env GOOS=$$GOOS GOARCH=$$GOARCH go build -o $(EXECUTABLES_DIR)/$$OUTPUT_NAME || exit 1; \
	done

setup:
	@rm -f standards-oui.ieee.org.txt
	@echo "Downloading OUI data..."
	@wget -q https://raw.githubusercontent.com/wbrijesh/network-scan/refs/heads/main/standards-oui.ieee.org.txt -O standards-oui.ieee.org.txt
	@if [ $$? -eq 0 ]; then \
		echo "✅ Successfully downloaded OUI data"; \
	else \
		echo "Error: Failed to download OUI data"; \
		exit 1; \
	fi

clean:
	@rm -rf $(EXECUTABLES_DIR)
	@rm -f standards-oui.ieee.org.txt
