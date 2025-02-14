import { useState } from "react";
import PageHeader from "../header/PageHeader";

function FlightCreate() {
    const [flight, setFlight] = useState({
        id: '1010',
        number: "AI 8499",
        airline_name: 'Air India',
        source: 'Mysore',
        destination: 'Trichy',
        capacity: '180',
        price: 9977.0
    });

    const onBoxChange = (event) => {
        const { id, value } = event.target;
        // Update state with a new object
        setFlight((prevFlight) => ({
            ...prevFlight,
            [id]: value
        }));
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        // You can handle the form submission logic here
        console.log(flight);
    };

    return (
        <>
            <PageHeader PageNumber={2} />
            <h3><a href="/flights/list" className="btn btn-light">Go Back</a>New Flight</h3>
            <div className="container">
                <form onSubmit={handleSubmit}>
                    <div className="form-group mb-3">
                        <label htmlFor="number" className="form-label">Flight Number:</label>
                        <input
                            type="text"
                            className="form-control"
                            id="number"
                            value={flight.number}
                            onChange={onBoxChange}
                            placeholder="Please enter flight number"
                        />
                    </div>
                    <div className="form-group mb-3">
                        <label htmlFor="airline_name" className="form-label">Airline Name:</label>
                        <input
                            type="text"
                            className="form-control"
                            id="airline_name"
                            value={flight.airline_name}
                            onChange={onBoxChange}
                            placeholder="Please enter airline name"
                        />
                    </div>
                    <div className="form-group mb-3">
                        <label htmlFor="source" className="form-label">Source:</label>
                        <input
                            type="text"
                            className="form-control"
                            id="source"
                            value={flight.source}
                            onChange={onBoxChange}
                            placeholder="Please enter source"
                        />
                    </div>
                    <div className="form-group mb-3">
                        <label htmlFor="destination" className="form-label">Destination:</label>
                        <input
                            type="text"
                            className="form-control"
                            id="destination"
                            value={flight.destination}
                            onChange={onBoxChange}
                            placeholder="Please enter destination"
                        />
                    </div>
                    <div className="form-group mb-3">
                        <label htmlFor="capacity" className="form-label">Capacity(Number of Seats):</label>
                        <input
                            type="text"
                            className="form-control"
                            id="capacity"
                            value={flight.capacity}
                            onChange={onBoxChange}
                            placeholder="Please enter capacity"
                        />
                    </div>
                    <div className="form-group mb-3">
                        <label htmlFor="price" className="form-label">Ticket Price:</label>
                        <input
                            type="text"
                            className="form-control"
                            id="price"
                            value={flight.price}
                            onChange={onBoxChange}
                            placeholder="Please enter ticket price"
                        />
                    </div>
                    <button type="submit" className="btn btn-success">Create Flight</button>
                </form>
            </div>
        </>
    );
}

export default FlightCreate;
