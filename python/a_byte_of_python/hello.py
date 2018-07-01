def hello(name: str, city: str):
    print("Hello! My name is {name}. I'm from {city}.".
          format(name=name, city=city), end='')


hello("ke", "shanghai")
