CREATE TABLE sid.client_ibu (
    docid character varying(75) UNIQUE,
    datecreated timestamp without time zone,
    baseentityid character varying(75),
    uniqueid character varying(75),
    namalengkap character varying(75),
    namasuami character varying(75),
    provinsi character varying(75),
    kabupaten character varying(75),
    kecamatan character varying(75),
    desa character varying(75),
    dusun character varying(75),
    birthdate date,
    nik character varying(75),
    noibu character varying(75),
    providerid character varying(75)
);


CREATE TABLE sid.client_anak (
    docid character varying(75) UNIQUE,
    datecreated timestamp without time zone,
    baseentityid character varying(75),
    uniqueid character varying(75),
    birthdate date,
    gender character varying(75),
    ibucaseid character varying(75),
    providerid character varying(75)
);

