# Saving files

>[!NOTE]
When saving data to files we need to save the data in byte slice format `[]byte` 

```go
os.WriteFile("bills/"+b.name+".txt", data, 0644)
```
when saving data to files we use the os package. inside the WriteFile method there is 3 arguments.

1. `location` were you want to save the file, `Name` you're assigning to the file and format `txt`
2. data
3. permissions (in the above example its `0644`)

>[!NOTE]
In Linux, file permissions are represented by a three-digit octal number. The permission 0644 means the following:

The first digit (0) indicates special modes (setuid, setgid, sticky bit), which are not set in this case.
The second digit (6) represents the owner's permissions:
4 (read)
2 (write)
4 + 2 = 6 (read and write)
The third digit (4) represents the group's permissions:
4 (read)
The fourth digit (4) represents the permissions for others (everyone else):
4 (read)
So, 0644 means:

The owner of the file has read and write permissions.
The group has read-only permissions.
Others (everyone else) have read-only permissions.

```go
panic(err)
```
panic is a method used to stop the flow of the program and print out an error.