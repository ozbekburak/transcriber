# Introduction:
As we know, transcribing audio files into written text is a time-consuming task. However, our new application will simplify this process by transcribing audio files and 
translating them into a user-specified language.

# Usage
To use the program, the user needs to set the [OPENAI API KEY](https://platform.openai.com/account/api-keys) environment variable to their API key for OpenAI. After that:
```go
    export OPENAI_API_KEY={your_api_key}
    git clone github.com/ozbekburak/transcriber
    cd transcriber
    go run main.go
```

or 

```go
    export OPENAI_API_KEY={your_api_key}
    git clone github.com/ozbekburak/transcriber
    cd transcriber
    go build .
    ./transcriber
```

# Screenshots
![execution](https://github.com/ozbekburak/transcriber/blob/main/img/execution.png?raw=true)
![files](https://github.com/ozbekburak/transcriber/blob/main/img/transcript-translation.png?raw=true)

## Features:

### Audio Transcriptions:
It transcribes audio files with the mp3/mp4 format (any format that supports chatgpt) into English. It uses speech recognition technology to convert the audio into text. The transcribed text is saved to a file under the directory name **transcriptions**.

### Language Translation:
It can translate the transcribed text into a specified language. The user can choose the language to be translated from a list of available languages. The application uses a translation engine to translate the text, and the translated text is saved to a file under the directory named **translations**.

### User-Friendly:
The user simply needs to specify the language in which they want the text translated when prompted. The application will take care of the rest. For example, if the user specifies "French" as the target language, the application will translate the text into French and save the file under the "translations" directory.

### File Management:
The application organizes the transcribed and translated files into separate directories. This makes it easy for the user to locate the files and use them for various purposes.

## Conclusion

- The application can be used by researchers to transcribe and translate audio files for various research purposes.

- The application can be used by businesses to transcribe and translate audio files for various business purposes such as creating transcripts of meetings and interviews, and translating them into different languages for international clients.

- It can also be useful for businesses that create videos on platforms such as YouTube and need a transcription of the audio to provide closed captions for their viewers.

It will significantly reduce the time and effort required for transcribing and translating audio files. This, in turn, will improve productivity and efficiency.
