# lib-cli - Command-Line Interface
lib-cli is the command-line client for the some generator APIs implement in this library! It provides simple access to all APIs functions to make a server, a GUI app & ...!

### Make new project - Use git as version control
- Make project folder and suggest use your domain name for it.
- Make project version control by `git init`
- Instead aboves you can clone exiting repo by `git clone ${repository path}`.
- Add libgo to project as submodule by `git submodule add -b master https://github.com/SabzCity/libgo`
- Build lib-cli by `go build ./libgo/lib-cli`
- Run lib-cli in a terminal by `./lib-cli`
- Choose desire services to make needed files or other actions!

### APIs
- Complete manifest in main package of service.
- Add other data to main package if needed.
- Add as many service you need by CLI services and add business logic to them!
- From CLI update service file to autogenerate some code for you!
- As you can see in file services logic layers are independent layer and you must just think locally! But if you need network stream data use `st *achaemenid.Stream` in your each function parameters. Don't remove it even don't need it!

## RUN
- first check and change `DevMode` and `DebugMode` const in log package to desire behavior
- Build app by `go build`
- Strongly suggest run app by systemd on linux or other app manager on other OS!
- Otherwise easily run app by `./{{root-folder-name}}`
