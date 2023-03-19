This is a short reproduction for what seems to be a Landlock bug

## Steps to reproduce

* Create a directory on a eCryptFS mount(!)
* Landlock the process, give yourself permission to read that directory
* Read the directory

## Expected behaviour

Should be able to read the directory

## Observed behaviour

Permission denied (`EACCES`)
