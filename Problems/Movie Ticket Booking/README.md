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
- Movie
- Theater
- Scree
- Seat
- Show
- ShowSeating
- Booking
- Ticket
- User

Relationships:

```text
    BookingSystem <---- contains --- movies, shows, theators, bookings

    Theator <---- contains ---- screens

    BookingSystem <---- contain ---- Movie
```

## Class Design

```code
Class BookingSystem

    - movies: Map<String, Movie>
    - theators: Map<String, Threator>
    - shows: Map<String, Show>
    - bookings: Map<String, Booking>

    + search(title: string) -> List<Movie>
    + getShow(id: string) Show
    + book(user, showID, seatIDs) BookingID
    + cancel(bookingID: string)
```

```code
Class Movie

    - id: string
    - title: string
    - description: string
    - startTime: time
```

```code
Class Threator:

    - id: stirng
    - name: string
    - screens: Map<String, Screen>

    + getScreen(id) -> List<Seat>
```

```code
Class Show

    - id: string
    - movieID: string
    - screen: Screen
    - startTime: time

    + reserveSeat(seatID: string)
    + availableSeats() -> List<Seat>
```

```code
Class Screen

    - id: string
    - seats: List<Seats>

    + reserveSeat(seatID: string)
    + availableSeats() -> List<Seat>
```

```code
Class Seat:

    - id: int
    - type: NORMAL, GOLD, PREMIUM
    - price: float
    - state: AVAILABEL | OCCUPIED

    + researSeat(id: string) 
```

```code
Class Booking

    - id: string
    - userID: string
    - showID: string
    - tickets: List<Ticket>
    - total: float
    - state: BOOKED | PENDING | CANCELED
```

```code
Class Ticket:

    - id: string
    - bookingID: string
    - showID: string
    - seatID: int
```
