## [singleheader-library-generator]
################################################################################
## BUILD
################################################################################
TARGET="singleheader-library-generator"
LIBRARY_SRC_PATH="./libraries"
################################################################################
LIBRARIES= cog 
################################################################################
# TODO: Prefix build with clean, and only clean if starship yard exists. 

all: $(LIBS)

build: 
		@echo ".==============================================."
		@echo "| [singleheader-library-generator]             |"
		@echo "'=============================================='"
		@echo "not implemented"
		@echo $(1)
		@echo "$(1) $2 $3 $4"

clean:
		# TODO: Only preform if bin/starship-yard exists
		# We want `clean all` vs. `clean cog`
		@echo ".==============================================."
		@echo "| [singleheader-library-generator]             |"
		@echo "'=============================================='"
		@echo "Cleaning up all generated libraries"
		@echo "$(1) $(2) $3 $4"

