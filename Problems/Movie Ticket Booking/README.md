# Movie Ticket Booking

Link: <https://www.hellointerview.com/learn/low-level-design/problem-breakdowns/bookmyshow>

## Clarifying Questions

- Seaching is full-text search, fuzzy-matching or just simple title lookup? [Ans: Simple title/text matching]
- Can user book multiple seat in one booking? [Ans: yes]
- Manage reservations include? Cancel, reschedule, modify? [Ans: Only Cancel]

Scope Questions:

- Are there different seat types with different prices? [Ans: Yes [NORMAL, GOLD, PREMIUM]]
- What about the concurrency? [Ans: Handle the concurrency]

## Requirements

- User should be able to search movies by title
- Theators have multipel screens; all screens share the same seat layout
- User should be able to view available seats
- User should be able to book the ticket
- User should be able to book multiple seats in single booking
- User should get ticket on the mail, phoneno
- The system should handle the concurrent bookings

Out Of Scope

- Ticket cancelation
- Payment Processing
- Notification System

## Entities & Relationships

- BookingSystem
- Booking
- Movie
- Theater
- User
- Ticket
- Seat

Relationships:

```text
    Theator <---- composed of ---- Seat

    BookingSystem <---- contain ---- Movie
```

## Class Design

```code
Class BookingSystem

    - movies: List<Movie>
    - theators: Map<String, Threator>
    - bookings: Map<String, Booking>

    + search(name: string) -> List<Movie>
    + book(booking: Booking) Ticket
```

```code
Class Movie

    - id: string
    - title: string
    - description: string
    - startTime: time
    - theatorID: string
```

```code
Class Theator:

    - id: stirng
    - seats: List<Seat>

    + getFreeSeat() -> List<Seat>
    + reserveSeat(Seat) 
```

```code
Class Seat:

    - id: string
    - type: NORMAL, GOLD, PREMIUM
    - price: float
    - state: AVAILABEL | OCCUPIED

    + researSeat(id: string) 
```

```code
Class Booking

    - id: string
    - tickets: List<Ticket>
    - price: float
    - state: BOOKED | PENDING | CANCELED
```

```code
Class Ticket:

    - id: string
    - movieID: string
    - seatID: string
```
