CREATE TABLE usd_rates (
    id SERIAL PRIMARY KEY,   
    rub FLOAT NOT NULL,    
    eur FLOAT NOT NULL,   
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE rub_rates (
    id SERIAL PRIMARY KEY,   
    usd FLOAT NOT NULL,    
    eur FLOAT NOT NULL,   
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE eur_rates (
    id SERIAL PRIMARY KEY,   
    usd FLOAT NOT NULL,    
    rub FLOAT NOT NULL,    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);



INSERT INTO usd_rates (rub, eur) VALUES
(101.23, 0.96);

INSERT INTO rub_rates (usd, eur) VALUES
(0.0099, 0.0095);

INSERT INTO eur_rates (usd, rub) VALUES
(1.04, 105.21);