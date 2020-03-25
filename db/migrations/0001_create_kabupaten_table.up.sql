CREATE TABLE kabupaten (
    id BIGSERIAL PRIMARY KEY,
    nama TEXT NOT NULL,
    odp INTEGER NOT NULL,
    pdp INTEGER NOT NULL,
    positif INTEGER NOT NULL,
    negatif INTEGER NOT NULL,
    meninggal INTEGER NOT NULL,
    selesai_pengawasan INTEGER NOT NULL,
    dalam_pengawasan INTEGER NOT NULL,
    selesai_pemantauan INTEGER NOT NULL,
    dalam_pemantauan INTEGER NOT NULL,
    updated_at TIMESTAMPTZ default now(),
    created_at TIMESTAMPTZ default now()
);
