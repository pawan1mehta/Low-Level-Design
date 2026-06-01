# SOLID Principles

SOLID is a set of five design principles. These principles help software developers design robust, testable, extensible, and maintainable object-oriented software systems.

Robert C. Martin introduced the SOLID principles in his 2000 paper “Design Principles and Design Patterns.” These concepts were later built upon by Michael Feathers, who introduced us to the SOLID acronym. In the last 20 years, these five principles have revolutionised the world of object-oriented programming, changing how we write software.

The SOLID principles are:
- **Single Responsibility Principle (SRP)**
- **Open/Closed Principle (OCP)**
- **Liskov Substitution Principle (LSP)**
- **Interface Segregation Principle (ISP)**
- **Dependency Inversion Principle (DIP)**


## Single Responsibility

This principle states that a class should only have one responsibility. Furthermore, it should only have one reason to change.


```
public class Book {
    private String name;
    private String author;
    private String text;
    //constructor, getters and setters
}
```

In this code, we store the name, author and text associated with an instance of a Book.
Let’s now add a couple of methods to query the text:

```
public class Book {
    private String name;
    private String author;
    private String text;

    //constructor, getters and setters

    // methods that directly relate to the book properties
    public String replaceWordInText(String word, String replacementWord){
        return text.replaceAll(word, replacementWord);
    }

    public boolean isWordInText(String word){
        return text.contains(word);
    }
}
```

Now our Book class works well, and we can store as many books as we like in our application.
But what good is storing the information if we can’t output the text to our console and read it?
Let’s throw caution to the wind and add a print method:

```
public class BadBook {
    //...
    void printTextToConsole(){
        // our code for formatting and printing the text
    }
}
```

However, this code violates the single responsibility principle we outlined earlier. To fix our mess, we should implement a separate class that deals only with printing our texts:

```
public class BookPrinter {
    // methods for outputting text
    void printTextToConsole(String text){
        //our code for formatting and printing the text
    }

    void printTextToAnotherMedium(String text){
        // code for writing to any other location..
    }
}
```

Awesome. Not only have we developed a class that relieves the Book of its printing duties, but we can also leverage our BookPrinter class to send our text to other media.
Whether it’s email, logging, or anything else, we have a separate class dedicated to this one concern.


## Open for Extension, Closed for Modification

Simply put, classes should be open for extension but closed for modification. In doing so, we stop ourselves from modifying existing code 
and causing potential new bugs in an otherwise happy application.


```
public class Guitar {
    private String make;
    private String model;
    private int volume;

    //Constructors, getters & setters
}
```

We launched the application, and everyone loves it. But after a few months, we decided the Guitar was a little boring and could use a cool flame pattern to make it look more rock and roll. At this point, it might be tempting to just open up the Guitar class and add a flame pattern — but who knows what errors might throw up in our application. Instead, let’s stick to the open-closed principle and simply extend our Guitar class:

```
public class SuperCoolGuitarWithFlames extends Guitar {
    private String flameColor;

    //constructor, getters + setters
}
```

By extending the Guitar class, we can be sure that our existing application won’t be affected.


## Liskov Substitution

If a program is using a base class, it should be able to use any of its subclasses without disrupting the behaviour of our program. It states that the objects of a subclass should behave the same way as the objects of the superclass, such that they are replaceable.


```
class Shape {
      public void draw() {
     	System.out.println("Drawing a shape");
      }
}

class Circle extends Shape {
    @Override
    public void draw() {
        System.out.println("Drawing a circle");
    }
}
class Square extends Shape {
    @Override
    public void draw() {
         System.out.println("Drawing a square");
    }
}
```

In this example, Circle and Square are subclasses of Shape. According to the Liskov Substitution Principle, you should be able to use instances of Circle and Square wherever you expect a Shape. 

```
public class Main {
    public static void drawShape(Shape shape) {
       shape.draw();
    }

     public static void main(String[] args) {
        Shape circle = new Circle();
        Shape square = new Square();

        drawShape(circle); // Output: Drawing a circle
        drawShape(square); // Output: Drawing a square
     }
}
```

Here, we're passing instances of Circle and Square to a method that expects a Shape, and it works perfectly fine. This demonstrates the Liskov Substitution Principle in action.


### Interface Segregation

The I in SOLID stands for interface segregation, and it simply means that larger interfaces should be split into smaller ones. By doing so, we can ensure that implementing classes only need to be concerned about the methods that are of interest to them.


```
public interface BearKeeper {
    void washTheBear();
    void feedTheBear();
    void petTheBear();
}
```

As avid zookeepers, we’re more than happy to wash and feed our beloved bears. But we’re all too aware of the dangers of petting them. Unfortunately, our interface is rather large, and we have no choice but to implement the code to pet the bear. Let’s fix this by splitting our large interface into three separate ones:
```
public interface BearCleaner {
    void washTheBear();
}

public interface BearFeeder {
    void feedTheBear();
}

public interface BearPetter {
    void petTheBear();
}


public class BearCarer implements BearCleaner, BearFeeder {

    public void washTheBear() {
        //I think we missed a spot...
    }

    public void feedTheBear() {
        //Tuna Tuesdays...
    }
}

public class CrazyPerson implements BearPetter {
    public void petTheBear() {
        //Good luck with that!
    }
}
```

## Dependency Inversion

It allows the creation of dependent objects outside of a class and provides those objects to a class in different ways.

High-level modules should not depend on low-level modules. Both should depend on the abstraction.
Abstractions should not depend on details. Details should depend on abstractions.

```
public class Windows98Machine {

}
```

But what good is a computer without a monitor and keyboard? Let’s add one of each to our constructor so that every Windows98Computer we instantiate comes prepacked with a Monitor and a StandardKeyboard:

```
public class Windows98Machine {
    private final StandardKeyboard keyboard;
    private final Monitor monitor;

    public Windows98Machine() {
        monitor = new Monitor();
        keyboard = new StandardKeyboard();
    }
}
```

This code will work, and we’ll be able to use the StandardKeyboard and Monitor freely within our Windows98Computer class.

Problem solved? Not quite. By declaring the StandardKeyboard and Monitor with the new keyword, we’ve tightly coupled these three classes together. Not only does this make our Windows98Computer hard to test, but we’ve also lost the ability to switch out our StandardKeyboard class with a different one should the need arise. And we’re stuck with our Monitor class too.

```
public interface Keyboard {

}

public class Windows98Machine{
    private final Keyboard keyboard;
    private final Monitor monitor;

    public Windows98Machine(Keyboard keyboard, Monitor monitor) {
        this.keyboard = keyboard;
        this.monitor = monitor;
    }
}
```

Here, we’re using the dependency injection pattern to facilitate adding the Keyboard dependency into the Windows98Machine class. Let’s also modify our StandardKeyboard class to implement the Keyboard interface so that it’s suitable for injecting into the Windows98Machine class:
```
public class StandardKeyboard implements Keyboard { 


}
```