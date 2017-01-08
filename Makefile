#Define Env variable for golang.
export GOVERSION="go1.7.1"
export GOPATH=$(PWD)

NAME = gomoku

BPKG = algo \
	   gomoku\
	   
SDL2 = github.com/veandco/go-sdl2/sdl

all: $(NAME)

$(NAME):
	@echo "\033[1;36;m[Compiling $@]\033[0m: "
	go build $(BPKG)
	go install $(NAME)
	@echo "\033[1;32;m[Done]\033[0m"

submodule:
	git submodule update --init --recursive 

clean:
	@rm -rf pkg
	@echo "Remove \033[1;31mPackage(s)\033[0m."

fclean: clean
	@rm -rf bin
	@echo "Remove \033[1;31mBin\033[0m."

re: fclean all
