#Define Env variable for golang.
export GOVERSION="go1.7.1"
export GOPATH=$(PWD)

NAME = gomoku
SDL2 = github.com/veandco/go-sdl2/sdl

all: $(NAME)

$(NAME):
	go get -v $(SDL2)
	go install $(NAME)
	@printf "Building \033[1;34m$(NAME)\033[0m."
	@sleep 0.7
	@printf "."
	@sleep 0.7
	@printf ".\n"


clean:
	@rm -rf pkg
	@echo "Remove \033[1;31mPackage(s)\033[0m."

fclean: clean
	@rm -rf bin
	@echo "Remove \033[1;31mBin\033[0m."

re: fclean all
