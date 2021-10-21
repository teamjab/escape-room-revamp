# Escape room revamp api

- This Golang project is to provide API backend with ability for user to create riddles and scores, and get riddles and scores. 

```json
// "/riddles" POST
{
    "Username":"PUT USERNAME",
    "Question":"SEND QUESTION",
    "Answer":[{"AnswerOne":"ANSWER ONE","AnswerTwo":"ANSWER TWO","AnswerThree":"ANSWER THREE","AnswerFour":"ANSWER FOUR"}],
    "CorrectAnswer": "SEND CORRECT ANSWER"
}

// "/socres" POST
{
    "Username": "PUT USERNAME",
    "Score": "PUT SOCRE IN INT TYPE"
}


// "/riddles" GET to get all scores
// "/scores" GET to get all scores
```

# Unit Test
- To run the test follow below

```sh
cd api
go test ./... -v # -v for verbose
```

# Local Development
- To run the application env file is needed in the `.env` file put following...
```txt
# .env

MONGODB_URI=<PUT URL FOR MONGO DB>
PORT=8080
```
- Dockerfile is available for local build
```sh
cd api

docker build -t api .

docker run api # add -d inbetween run and api if you want detached mode
```
