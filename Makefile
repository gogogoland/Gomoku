#Define Env variable for golang.
export GOVERSION="go1.7.1"
export GOPATH=$(PWD)

NAME = gomoku
BPKG = algo\
	   gomoku\
	   
SDL2 = github.com/veandco/go-sdl2/sdl

all: $(NAME)

$(NAME):
	go get -v $(SDL2)
	go build $(BPKG)
	go install $(NAME)
	@printf "Building \033[1;34m$(NAME)\033[0m."

submodule:
	git submodule update --init --recursive 

clean:
	@rm -rf pkg
	@echo "Remove \033[1;31mPackage(s)\033[0m."

fclean: clean
	@rm -rf bin
	@echo "Remove \033[1;31mBin\033[0m."

re: fclean all
