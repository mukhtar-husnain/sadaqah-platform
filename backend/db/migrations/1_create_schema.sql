-- Users Table (NGOs, madrassas, individuals, etc.)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,  -- Hashed password
    user_type VARCHAR(50) CHECK (user_type IN ('NGO', 'Individual', 'Madrassa', 'Charity Group')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Volunteer Requests Table (Help posts)
CREATE TABLE volunteer_requests (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    location VARCHAR(255),
    event_date TIMESTAMP NOT NULL,
    required_volunteers INT NOT NULL,
    skills_required TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Volunteers Table (Users who volunteer for requests)
CREATE TABLE volunteers (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    request_id INT REFERENCES volunteer_requests(id) ON DELETE CASCADE,
    status VARCHAR(50) CHECK (status IN ('Pending', 'Accepted', 'Rejected')) DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Feedback Table (Post-event ratings and reviews)
CREATE TABLE feedback (
    id SERIAL PRIMARY KEY,
    request_id INT REFERENCES volunteer_requests(id) ON DELETE CASCADE,
    given_by_user_id INT REFERENCES users(id) ON DELETE CASCADE,
    given_to_user_id INT REFERENCES users(id) ON DELETE CASCADE,
    rating INT CHECK (rating BETWEEN 1 AND 5),
    comments TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
