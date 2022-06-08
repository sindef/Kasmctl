Kasmctl is just intended to be an easy CLI tool to manage our kasm environment. It's quite hastily thrown together, and I am not a professional developer (hence the code quality!)

You need to either set the environmental variables for:

```
KASM_URL
KASM_SECRET
KASM_KEY
```

or place these in a config file ($HOME/.kasmctl/config default) - if this file does not exist, it will be written with the environmental variables currently set
