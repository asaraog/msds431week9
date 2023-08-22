# Chatbot Prototype Application

## Project Summary
This project aims to produce a prototype for a desktop chatbot application based on [Wails](https://wails.io/) and [Svelte](https://svelte.dev/repl/hello-world). It aims to accept plain text input and output answers from a knowledge base represented with [Go's SQL](https://github.com/mattn/go-sqlite3) driver. For this prototype, only one word inputs will be allowed and a [simple corpus](./QandA.csv) of possible question/answer pairs will be used. The prototype is capable of integrating NLP concepts such as [term frequency-inverse document frequency (TF-IDF)](https://yi-wang-2005.medium.com/nlp-in-sql-word-vectors-82dffc908423). It also includes support for errors such as too many words, no input and no matches. The application prototype is succesful during development and build in providing the correct answers to user input questions.

Go should suffice for most software components (database, frontend, backend) and can be used for language processing as well with [wego](https://github.com/ynqa/wego) package for word embeddings and [spago](https://github.com/nlpodyssey/spago) for neural networks and LLMs like huggingface. The SQL version of the database is a great start but using a hosted database of vector objects like [Milvus](https://milvus.io/) is recommended for the final project as it would be easier to integrate with OpenAI's embedding representations.

## Further development as an AI application
Further development of the application would first continue in SQL by adding a TF-IDF representation to calculate similarity scores. We can then add 'fuzzy matching' and tokenization/lemmatization of the corpus and query. This will have the simplest development process with no dependency required. The single word query is highly amenable to using word2vec.

The next step would be to add [embeddings from OpenAI](https://www.singlestore.com/blog/getting-openai-embeddings-in-sql-using-external-functions/) which will require an external dependence.  We would then add LLM support using huggingface. This might be complicated with Wails based on previous experience.

## Important files

**app.go:** Backend 'brain' of application. Lookup function creates an SQL database for query. A test unit function is written in app_test.go.

**frontend/src/App.Svelte:** Frontend specifications. Binding of user input for preference to backend code. Displays output of Lookup.

**./build/bin/Week9.app** Executable application for MacOS.

**QandA.csv:** Corpus of questions/answers from [others](https://github.com/ThomasWMiller/jump-start-sqlite/blob/main/QandA.csv).

## Installation and Running

First install Wails onto your machine.
```
xcode-select --install
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

Download or git clone this project onto your local machine and test using vale CLI before running the development application:
```
git clone https://github.com/asaraog/msds431week9.git
cd msds431week9/build/bin
open Week9.app
```
Before clicking 'Query', input nothing, 'break', 'test' and 'break test' in the text box to check for correct application behavior.
