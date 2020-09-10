package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	doc = []string {
	    "FROM ruby:2.5\n",
		"RUN apt-get update -qq && apt-get install -y nodejs postgresql-client\n",
		"RUN mkdir /myapp\n",
		"WORKDIR /myapp\n",
		"COPY Gemfile /myapp/Gemfile\n",
		"COPY Gemfile.lock /myapp/Gemfile.lock\n",
		"RUN bundle install\n",
		"COPY . /myapp\n\n",
		"# Add a script to be executed every time the container starts.\n",
		"COPY entrypoint.sh /usr/bin/\n",
		"RUN chmod +x /usr/bin/entrypoint.sh\n",
		"ENTRYPOINT [\"entrypoint.sh\"]\n",
		"EXPOSE 3000\n\n",
		"# Start the main process.\n",
		"CMD [\"rails\", \"server\", \"-b\", \"0.0.0.0\"]",
    }
	composeDoc = []string{
		"version: '3'\n",
		"services:\n",
		"  db:\n",
		"    image: postgres\n",
		"    volumes:\n",
		"      - ./tmp/db:/var/lib/postgresql/data\n",
		"    environment:\n",
		"      POSTGRES_PASSWORD: password\n",
		"  web:\n",
		"    build: .\n",
		"    command: bash -c \"rm -f tmp/pids/server.pid && bundle exec rails s -p 3000 -b '0.0.0.0'\"\n",
		"    volumes:\n",
		"      - .:/myapp\n",
		"    ports:\n",
		"      - \"3000:3000\"\n",
		"    depends_on:\n",
		"      - db",
	}
)

func main() {
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	path := args[1] + "/"
	file, err := os.OpenFile(path + "Dockerfile", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	comFile, err := os.OpenFile(path + "docker-compose.yml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer comFile.Close()
	_ = writeByres(file, doc)
	_ = writeByres(comFile, composeDoc)
	out, err := exec.Command("ls", "-la").Output()
	fmt.Println(string(out))
}

func writeByres(file *os.File, array []string) error {
	for _, line := range array {
		b := []byte(line)
		_, err := file.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}
