# Demos for the WASM talk

```bash
./emsdk activate latest
source ~/dev/emsdk/emsdk_env.sh
```

* [Skelets](http://aws-website-webassemblyskeletalanimation-ffaza.s3-website-us-east-1.amazonaws.com/)
* [QT](http://example.qt.io/qt-webassembly/widgets/richtext/textedit/textedit.html )
* [sum](https://mbebenita.github.io/WasmExplorer/?state=%7B%22options%22%3A%7B%22showGutter%22%3Atrue%2C%22showConsole%22%3Atrue%2C%22showOptions%22%3Atrue%2C%22autoCompile%22%3Atrue%2C%22showLLVM%22%3Afalse%2C%22darkMode%22%3Atrue%2C%22fastMath%22%3Afalse%2C%22noInline%22%3Afalse%2C%22noRTTI%22%3Afalse%2C%22noExceptions%22%3Afalse%2C%22cleanWast%22%3Afalse%2C%22wasmBaseline%22%3Afalse%2C%22dialect%22%3A%22C99%22%2C%22optimizationLevel%22%3A%22s%22%7D%2C%22source%22%3A%22int%20sum(int%20a%2C%20int%20b)%7B%5Cn%20%20return%20a%20%2B%20b%3B%5Cn%7D%22%7D)
* [alert](https://mbebenita.github.io/WasmExplorer/?state=%7B%22options%22%3A%7B%22showGutter%22%3Atrue%2C%22showConsole%22%3Atrue%2C%22showOptions%22%3Atrue%2C%22autoCompile%22%3Atrue%2C%22showLLVM%22%3Afalse%2C%22darkMode%22%3Atrue%2C%22fastMath%22%3Afalse%2C%22noInline%22%3Afalse%2C%22noRTTI%22%3Afalse%2C%22noExceptions%22%3Afalse%2C%22cleanWast%22%3Afalse%2C%22wasmBaseline%22%3Afalse%2C%22dialect%22%3A%22C99%22%2C%22optimizationLevel%22%3A%22s%22%7D%2C%22source%22%3A%22void%20alert(int%20val)%3B%5Cn%5Cnvoid%20sum(int%20a%2C%20int%20b)%7B%5Cn%20%20int%20res%20%3D%20a%20%2B%20b%3B%5Cn%20%20alert(res)%3B%5Cn%7D%22%7D)
* [memory](https://mbebenita.github.io/WasmExplorer/?state=%7B%22options%22%3A%7B%22showGutter%22%3Atrue%2C%22showConsole%22%3Atrue%2C%22showOptions%22%3Atrue%2C%22autoCompile%22%3Atrue%2C%22showLLVM%22%3Afalse%2C%22darkMode%22%3Atrue%2C%22fastMath%22%3Afalse%2C%22noInline%22%3Afalse%2C%22noRTTI%22%3Afalse%2C%22noExceptions%22%3Afalse%2C%22cleanWast%22%3Afalse%2C%22wasmBaseline%22%3Afalse%2C%22dialect%22%3A%22C99%22%2C%22optimizationLevel%22%3A%22s%22%7D%2C%22source%22%3A%22void%20alert(char*%20str)%3B%5Cn%5Cnvoid%20doWork()%7B%5Cn%20%20%20%20alert(%5C%22Hello%20from%20WASM%5C%22)%3B%5Cn%7D%5Cn%22%7D)
