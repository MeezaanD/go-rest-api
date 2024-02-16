import { useState, useEffect } from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import { Home } from './Home.tsx'


interface Vehicle {
	id: string;
	brand: string;
	Name: string;
	type: string;
	image: string;
	price: number;
}

function App() {
	const [vehicles, setVehicles] = useState<Vehicle[]>([]);

	useEffect(() => {
		fetch('/api/vehicles')
			.then(response => response.json())
			.then(data => {
				console.log('API Response:', data);
				if (data.status === 'success') {
					setVehicles(data.vehicles);
				} else {
					console.error('Error fetching data:', data.message);
				}
			})
			.catch(error => console.error('Error fetching data:', error));
	}, []);

	return (
		<>
			<section className="bg-body-tertiary" id="hero">
				<Home />
				<div className="container-fluid w-100">
					<div className="row d-flex justify-content-center align-items-center flex-wrap">
						{vehicles.map(vehicle => (
							<div key={vehicle.id} className="card shadow rounded-3 m-2" style={{ maxWidth: '27rem' }}>
								<img src={vehicle.image} style={{width: ''}} alt={vehicle.Name} />
								<div className="card-body">
									<h5 className="card-title text-start">{vehicle.brand} - {vehicle.Name}</h5>
									<p className="card-text text-start">Type: {vehicle.type}</p>
									<p className="card-text text-start">Price: R{vehicle.price}</p>
								</div>
							</div>
						))}
					</div>
				</div>

			</section>
		</>
	);
}

export default App;
