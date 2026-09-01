# Amanzon Locker

Link: <https://www.hellointerview.com/learn/low-level-design/problem-breakdowns/amazon-locker>

## Clarifying Questions

- Are there different sized compartments? [ans: Yes small, medium, large]
- What's the scope of this system?
- Can one customer have multiple packages in the system at once? Are access tokens unique per package?
- How delivery driver will deposite a package into available compartant?
- Will there be infinite number of compartants?

## Requirements

- Carrier should be able to deposite a package by specifying the small, medium or large size
  - upon successfull deposite an access token should be generated send to user
- User should be able to retrive the package by entering the access token
  - system validate the code and opens the compartmant
  - throws error if the code is invalid or expired
- Access token expire after 7 days

Out Of Scope

- How package gets to the locker (delivery)
- How access token reaches the customer
- UI layer
- Payment

## Entities & Relationships

- Locker
- Compartment
- AccessToken

Relationships:

Locker <----composed of ---- Compartment

## Class Design

```code
Class Locker

- compartments: []Compartments
- accessTokenMap: Map<string, AccessToken>

+ Locker(compartment)
+ depositePackage(size) -> string | error 
+ pickup(tokenCode) -> void | error
+ openExpiredCompartments()
```

```code
Class AccessToken

- code: int
- expiration: timestamp
- compartment: Compartment

+ isExpired() -> bool
+ getCompartment() -> Compartment
+ getCode() -> string
```

```code
Class Compartment

- size: Size
- occupied: bool

+ isOccupied() -> bool
+ marOccupied()
+ getSize() -> Size
+ markFree()
```
