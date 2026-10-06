#include <iostream>

using namespace std;

class Shape {
public:
    virtual ~Shape() = default;
    virtual void draw() const = 0;
};

class Rectangle : public Shape {
public:
    void draw() const override {
        cout << "Drawing Rectangle ..." << endl;
    }
};

class Circle : public Shape {
public:
    void draw() const override {
        cout << "Drawing Circle ..." << endl;
    }
};

class ShapeFactory {
public:
    static unique_ptr<Shape> create(const string& type) {
        if (type == "circle") {
            return make_unique<Circle>();
        } else if (type == "rectangle") {
            return make_unique<Rectangle>();
        }
        throw invalid_argument("Unknown shape type: " + type);
    }
};

int main() {
    const unique_ptr<Shape> circleShape = ShapeFactory::create("circle");
    circleShape->draw();
    return 0;
}