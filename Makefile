#Define Env variable for golang.
export GOVERSION="go1.7.3"
export GOPATH=$(PWD)

NAME = gomoku

BPKG = algo \
	   gui \
	   gomoku \
	   
all: $(NAME)

$(NAME): library
	@echo "\033[1;36;m[Compiling $@]\033[0m: "
	go build $(BPKG)
	go install $(NAME)
	@echo "\033[1;32;m[Done]\033[0m"

library:
	@if [ ! -d "$(PWD)/src/golang.org" ] && [ ! -d "$(PWD)/src/github.com" ];then \
		echo "\033[1;36;m[Getting $@...]\033[0m "; \
		go get -u github.com/google/gxui/...; \
	fi

clean:
	@rm -rf pkg
	@echo "Remove \033[1;31mPackage(s)\033[0m."

fclean: clean
	@rm -rf bin
	@echo "Remove \033[1;31mBin\033[0m."

re: fclean all
