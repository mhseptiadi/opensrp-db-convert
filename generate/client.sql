CREATE TABLE sid2.client_ibu (
    docId character varying(75),
    dateCreated date,
    baseEntityId character varying(75),
    uniqueId character varying(75),
    namaLengkap character varying(75),
    namaSuami character varying(75),
    provinsi character varying(75),
    kabupaten character varying(75),
    kecamatan character varying(75),
    desa character varying(75),
    dusun character varying(75),
    birthDate date,
    nik character varying(75),
    noIbu character varying(75),
    providerId character varying(75),
);

CREATE TABLE sid2.client_anak (
    docId character varying(75),
    dateCreated date,
    baseEntityId character varying(75),
    uniqueId character varying(75),
    birthDate date,
    gender character varying(75),
    ibuCaseId character varying(75),
    providerId character varying(75),
);