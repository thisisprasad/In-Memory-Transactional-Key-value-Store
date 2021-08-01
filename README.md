# Simple Key/Value store in golang

## Store
An interactive shell based transactional in-memory key/value store inspired from here. [here](https://www.freecodecamp.org/news/design-a-key-value-store-in-go/).

An example script.
```sql 
> set val1 30
> set val2 55
> begin
> set val1 100
> get val1
100
> commit
> begin
> set val2 200
> get val2
200
> end
> get val2
55
> get val1
100
```

## How to use
1. Clone the project
2. Open the folder in command prompt
3. Run "go build"
4. Run "KeyValueStore" to run the shell.