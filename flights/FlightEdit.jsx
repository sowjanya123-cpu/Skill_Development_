import { useEffect, useState } from "react";
import PageHeader from "../header/PageHeader";
import { useNavigate, useParams } from "react-router-dom";
import axios from "axios";

function FlightEdit() {
    const [flight, setFlight] = useState({id:'', number:'', airline_name:'',
                source:'', destination:'', capacity: 0, price: 0.0})
    const OnBoxChange = (event) => {
        const newFlight = {...flight};
        newFlight[event.target.id] = event.target.value;
        setFlight(newFlight);
    }
    const params = useParams();
    const readFlightById = async () => {
        alert(params.id);
        try {
            const baseUrl = 'http://localhost:8080'
            const response = await axios.get(`${baseUrl}/flights/${params.id}`);
            setFlight(response.data);
            
        } catch(error) {
            alert('Server Error');
        }
    };
    useEffect(()=>{ readFlightById(); },[]);
    const navigate = useNavigate();
    const OnUpdate = async () => {
        try {
            const baseUrl = 'http://localhost:8080'
            const response = await axios.put(`${baseUrl}/flights/${params.id}`, {...flight,
                                        price:parseFloat(flight.price)});
            alert(response.data.message)
            navigate('/flights/list');
        } catch(error) {
            alert('Server Error');
        }
    }
    return (
        <>
            <PageHeader  PageNumber={1}/>
            <h3><a href="/flights/list" className="btn btn-light">Go Back</a>Edit Flight Ticket Price</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="number" className="form-label">Flight Number:</label>
                    <div className="form-control" id="number">{flight.number}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="airline_name" className="form-label">Airline Name:</label>
                    <div type="text" className="form-control" id="airline_name">{flight.airline_name}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="source" className="form-label">Source:</label>
                    <div className="form-control" id="source">{flight.source}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="destination" className="form-label">Destination:</label>
                    <div className="form-control" id="destination">{flight.destination}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="capacity" className="form-label">Capacity(Number of Seats):</label>
                    <div className="form-control" id="capacity">{flight.capacity}</div>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="price" className="form-label">Ticket Price:</label>
                    <input type="text" className="form-control" id="price" 
                        placeholder="Please enter ticket price"                                                
                        value={flight.price} onChange={OnBoxChange} />
                </div>
                <button className="btn btn-warning"
                    onClick={OnUpdate}>Update Price</button>
            </div>
        </>
    );
}

export default FlightEdit;