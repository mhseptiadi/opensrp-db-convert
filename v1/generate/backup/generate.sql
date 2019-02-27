CREATE TABLE IF NOT EXISTS sid.penutupan_ibu (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		otherclosereason character varying(75),
sub_village character varying(75),
cityvillage character varying(75),
address1 character varying(75),
closereason character varying(75),
iskicloseconfirmed character varying(75));



CREATE TABLE IF NOT EXISTS sid.update_photo (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		address1 character varying(75),
face_vector character varying(75),
cityvillage character varying(75));



CREATE TABLE IF NOT EXISTS sid.edit_tambah_kb (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		cityvillage character varying(75),
province character varying(75),
address1 character varying(75),
tanggalkunjungan character varying(75),
keterangantentangpesertakb character varying(75),
jeniskontrasepsi character varying(75),
district character varying(75),
tanggalperiksa character varying(75),
lokasiperiksa character varying(75),
tanggallahir character varying(75),
umur character varying(75),
paritas character varying(75));



CREATE TABLE IF NOT EXISTS sid.tambah_bayi (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		sub_village character varying(75),
deviceid character varying(75),
address1 character varying(75),
lokasiperiksa character varying(75),
tanggalpendaftaran character varying(75),
instanceid character varying(75),
simserial character varying(75),
phonenumber character varying(75),
cityvillage character varying(75),
tanggalperiksa character varying(75));



CREATE TABLE IF NOT EXISTS sid.penutupan_kb (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		address1 character varying(75),
kbclosereason character varying(75),
iskbcloseconfirmed character varying(75),
jeniskontrasepsi character varying(75),
cityvillage character varying(75));



CREATE TABLE IF NOT EXISTS sid.rencana_persalinan (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		ancid character varying(75),
address1 character varying(75),
lokasiperiksa character varying(75),
tempatrencanapersalinan character varying(75),
transportasi character varying(75),
pendonor character varying(75),
kondisirumah character varying(75),
cityvillage character varying(75),
htp character varying(75),
tanggalkunjunganrencanapersalinan character varying(75),
rencanapenolongpersalinan character varying(75),
rencanapendampingpersalinan character varying(75));



CREATE TABLE IF NOT EXISTS sid.kunjungan_anc_integrasi (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		high_risk_tbc character varying(75),
address1 character varying(75),
lokasiperiksa character varying(75),
referencedate character varying(75),
integrasiprogrampmtctvct character varying(75),
integrasiprogrampmtctperiksadarah character varying(75),
integrasiprogrampmtctserologi character varying(75),
integrasiprogrampmtctarvprofilaksis character varying(75),
integrasiprogrammalariaobat character varying(75),
cityvillage character varying(75),
integrasiprogramtbdahak character varying(75),
high_risk_malaria character varying(75),
ancid character varying(75),
integrasiprogrammalariakelambuberinsektisida character varying(75),
integrasiprogrammalariaperiksadarah character varying(75));



CREATE TABLE IF NOT EXISTS sid.dokumentasi_persalinan (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		tanggalkalaii character varying(75),
highriskpostpartumhemorrhage character varying(75),
sub_village character varying(75),
komplikasi character varying(75),
highriskpostpartumpih character varying(75),
highriskpostpartuminfection character varying(75),
jamplasentalahir character varying(75),
treatment character varying(75),
persentasi character varying(75),
integrasiprogram character varying(75),
penolong character varying(75),
jamkalaiaktif character varying(75),
tanggalplasentalahir character varying(75),
perdarahankalaiv2jampostpartum character varying(75),
keadaanbayi character varying(75),
highriskpostpartumdistosia character varying(75),
highriskpostpartumvacum character varying(75),
cityvillage character varying(75),
address1 character varying(75),
ancid character varying(75),
highriskpostpartumpreeclampsiaeclampsia character varying(75),
highriskpostpartummaternalsepsis character varying(75),
alamatbersalin character varying(75),
tanggalkalaiaktif character varying(75),
persalinan character varying(75),
jamkalaii character varying(75),
carapersalinanibu character varying(75),
highriskpostpartumforceps character varying(75),
tempatbersalin character varying(75),
dirujukke character varying(75),
tanggalhpht character varying(75),
keadaanibu character varying(75),
manajemenaktifkalaiii character varying(75));



CREATE TABLE IF NOT EXISTS sid.kohort_pelayanan_kb (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		keterangantentangpesertakb character varying(75),
jeniskontrasepsi character varying(75),
keterangantentangpesertakb2 character varying(75),
keteranganganticara character varying(75),
sub_village character varying(75),
cityvillage character varying(75),
address1 character varying(75),
tanggalkunjungan character varying(75));



CREATE TABLE IF NOT EXISTS sid.kunjungan_balita (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		panjangbayi character varying(75),
hasildilakukannyakpsp character varying(75),
pelayananvita character varying(75),
rujukanbayi character varying(75),
tanggalkunjunganbayiperbulan character varying(75),
lokasiperiksa character varying(75),
babyage character varying(75),
beratbayi character varying(75),
indikatorberatbadanbayi character varying(75),
statussesehatanbayi character varying(75),
mtbs character varying(75),
penyakit character varying(75),
desa_anak character varying(75));



CREATE TABLE IF NOT EXISTS sid.tambah_kb (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		tanggalkunjungan character varying(75),
lokasiperiksa character varying(75),
paritas character varying(75),
jeniskontrasepsi character varying(75),
tanggalperiksa character varying(75),
sub_district character varying(75),
cityvillage character varying(75),
tanggallahir character varying(75),
tdsistolik character varying(75),
alkilila character varying(75),
address1 character varying(75),
alkipenyakitkronis character varying(75),
keterangantentangpesertakb character varying(75),
province character varying(75),
district character varying(75),
umur character varying(75),
tddiastolik character varying(75),
alkihb character varying(75));



CREATE TABLE IF NOT EXISTS sid.identitas_ibu (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		agama character varying(75),
nomortelponhp character varying(75),
pasienwilayah character varying(75),
pasienpindahan character varying(75),
village character varying(75),
district character varying(75),
posyandu character varying(75),
jeniskontrasepsi character varying(75),
gakintidak character varying(75),
lokasiperiksaother character varying(75),
province character varying(75),
sub_district character varying(75),
puskesmas character varying(75),
pekerjaan character varying(75),
sub_village character varying(75),
asuransijiwa character varying(75),
lokasiperiksa character varying(75),
umur character varying(75),
pendidikan character varying(75),
golongandarah character varying(75),
registrationdate character varying(75),
tglpindah character varying(75));



CREATE TABLE IF NOT EXISTS sid.kunjungan_neonatal (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		kunjunganneonatal character varying(75),
beratbayi character varying(75),
kondisibayi character varying(75),
highriskchilditerus character varying(75),
highriskchildbirthtrauma character varying(75),
highriskchildinfection character varying(75),
sub_village character varying(75),
hb0 character varying(75),
komplikasineonatal character varying(75),
highriskchildcongenitalabnormality character varying(75),
highriskchildneonatalsepsis character varying(75),
suhubayi character varying(75),
frekuensinapas character varying(75),
frekuensidenyutjantung character varying(75),
highriskchildtetanusneonatorum character varying(75),
desa_anak character varying(75),
tanggalkunjunganbayiperbulan character varying(75),
jenispelayanan character varying(75),
panjangbayi character varying(75),
highriskchildhypothermia character varying(75),
highriskchildasphyxia character varying(75));



CREATE TABLE IF NOT EXISTS sid.post_partum_kb (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		metodekontrasepsi character varying(75),
rencana character varying(75),
ancid character varying(75),
cityvillage character varying(75),
address1 character varying(75));



CREATE TABLE IF NOT EXISTS sid.penutupan_pnc (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		cityvillage character varying(75),
address1 character varying(75),
pncclosereason character varying(75),
ancid character varying(75));



CREATE TABLE IF NOT EXISTS sid.edit_bayi (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		namabayi character varying(75),
desa_anak character varying(75),
beratlahir character varying(75));



CREATE TABLE IF NOT EXISTS sid.edit_ibu (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		district character varying(75),
pekerjaan character varying(75),
posyandu character varying(75),
golongandarah character varying(75),
jeniskontrasepsi character varying(75),
nomortelponhp character varying(75),
lokasiperiksaother character varying(75),
pasienwilayah character varying(75),
province character varying(75),
umur character varying(75),
registrationdate character varying(75),
pendidikan character varying(75),
tglpindah character varying(75),
asuransijiwa character varying(75),
sub_district character varying(75),
sub_village character varying(75),
agama character varying(75),
gakintidak character varying(75),
lokasiperiksa character varying(75),
pasienpindahan character varying(75),
village character varying(75),
puskesmas character varying(75));



CREATE TABLE IF NOT EXISTS sid.tambah_anc (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		highriskheartdisorder character varying(75),
lokasiperiksaother character varying(75),
tanggalhpht character varying(75),
ancid character varying(75),
hidup character varying(75),
tbcm character varying(75),
highriskpregnancytoomanychildren character varying(75),
pptest character varying(75),
highriskpregnancyproteinenergymalnutrition character varying(75),
address1 character varying(75),
trimesterke character varying(75),
htp character varying(75),
usiaklinis character varying(75),
highrisklaboursectioncesarearecord character varying(75),
highriskdidneydisorder character varying(75),
highriskmalaria character varying(75),
type character varying(75),
persalinansebelumnya character varying(75),
referencedate character varying(75),
jamkesmas character varying(75),
penyakitkronis character varying(75),
alergi character varying(75),
province character varying(75),
gravida character varying(75),
highriskpregnancyabortus character varying(75),
highriskectopicpregnancy character varying(75),
highrisktuberculosis character varying(75),
highriskhivaids character varying(75),
lokasiperiksa character varying(75),
bbkg character varying(75),
sub_district character varying(75),
riwayatkomplikasikebidanan character varying(75),
district character varying(75),
partus character varying(75),
tanggalmemperolehbukukia character varying(75),
tanggallahiranaksebelumnya character varying(75),
cityvillage character varying(75),
hasilpemeriksaanlila character varying(75),
highrisklabourtbrisk character varying(75),
highriskcardiovasculardiseaserecord character varying(75),
highriskasthma character varying(75),
abortus character varying(75));



CREATE TABLE IF NOT EXISTS sid.kohort_kunjungan_bayi_perbulan (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		statusgizi character varying(75),
hasildilakukannyakpsp character varying(75),
pelayananvita character varying(75),
babyage character varying(75),
kunjunganbayike character varying(75),
indikatorberatbadanbayi character varying(75),
asiaksklusif character varying(75),
desa_anak character varying(75),
tanggalkunjunganbayiperbulan character varying(75),
lokasiperiksa character varying(75),
statuskesehatanbayi character varying(75),
mtbs character varying(75),
rujukanbayi character varying(75),
sub_village character varying(75),
beratbayi character varying(75),
panjangbayi character varying(75));



CREATE TABLE IF NOT EXISTS sid.penutupan_anc (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		closereason character varying(75),
referred character varying(75),
referrallocation character varying(75),
birthattendant character varying(75),
otherclosereason character varying(75),
cityvillage character varying(75),
sub_village character varying(75));



CREATE TABLE IF NOT EXISTS sid.kunjungan_anc_lab_test (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		lokasiperiksa character varying(75),
laboratoriumperiksahbanemia character varying(75),
laboratoriumguladarah character varying(75),
treatmentanemiatxt character varying(75),
treatmenthbsagtxt character varying(75),
laboratoriumsifilis character varying(75),
cityvillage character varying(75),
laboratoriumperiksahbhasil character varying(75),
laboratoriumhbsag character varying(75),
highriskpregnancyanemia character varying(75),
highriskpregnancydiabetes character varying(75),
laboratoriumproteinuria character varying(75),
ancid character varying(75),
address1 character varying(75),
referencedate character varying(75),
laboratoriumthalasemia character varying(75));



CREATE TABLE IF NOT EXISTS sid.penutupan_anak (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		closereason character varying(75),
childdeathcause character varying(75),
placeofdeath character varying(75),
prereferralmanagement character varying(75),
referred character varying(75),
referrallocation character varying(75),
desa_anak character varying(75),
existing_sub_village character varying(75));



CREATE TABLE IF NOT EXISTS sid.child_registration (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		namabayi character varying(75),
desa_anak character varying(75),
beratlahir character varying(75),
highriskchildlowbirthweght character varying(75));



CREATE TABLE IF NOT EXISTS sid.kunjungan_anc (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		cityvillage character varying(75),
usiaklinis character varying(75),
ancke character varying(75),
tandavitaltddiastolik character varying(75),
tfu character varying(75),
taksiranberatjanin character varying(75),
highrisklabourfetusmalpresentation character varying(75),
ancid character varying(75),
dirujukke character varying(75),
jumlahjanin character varying(75),
kepalajaninterhadappap character varying(75),
highriskpregnancyproteinenergymalnutrition character varying(75),
highriskpregnancypih character varying(75),
pelayananfe character varying(75),
nomortelpon character varying(75),
caramasuktempatpelayanan character varying(75),
keadaanibusaatpulang character varying(75),
kunjunganke character varying(75),
district character varying(75),
sub_district character varying(75),
pelayanancatatdibukukia character varying(75),
province character varying(75),
pelayananinjeksitt character varying(75),
persentasijanin character varying(75),
statusimunisasitt character varying(75),
trimesterke character varying(75),
djj character varying(75),
nomorhp character varying(75),
address1 character varying(75),
komplikasidalamkehamilan character varying(75),
komplikasilain character varying(75),
keadaanibusaattiba character varying(75),
ancdate character varying(75),
tandavitaltdsistolik character varying(75),
highrisklabourfetusnumber character varying(75),
highrisklabourfetussize character varying(75),
hiddenancke character varying(75),
hasilpemeriksaanlila character varying(75),
keterangank1k4who character varying(75),
jamkesmas character varying(75),
lokasiperiksa character varying(75),
treatmenttext character varying(75),
treatment character varying(75),
bbkg character varying(75),
statusgiziibu character varying(75),
reflekspatelaibu character varying(75),
sub_village character varying(75),
anamnesisibu character varying(75));



CREATE TABLE IF NOT EXISTS sid.imunisasi_bayi (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		dpt_hb_lanjutan character varying(75),
polio4 character varying(75),
dpthb3 character varying(75),
polio1 character varying(75),
dpthb1 character varying(75),
campak character varying(75),
ipv character varying(75),
bcg character varying(75),
sub_village character varying(75),
hb0 character varying(75),
polio2 character varying(75),
dpthb2 character varying(75),
polio3 character varying(75),
desa_anak character varying(75));



CREATE TABLE IF NOT EXISTS sid.kunjungan_pnc (
		first_name character varying(75),
		last_name character varying(75),
		birth_date character varying(75),
		
		event_id character varying(75) UNIQUE,
		document_id character varying(75),
		base_entity_id character varying(75),	
		location_id character varying(75),	
		
		date_created timestamp without time zone,
		event_date timestamp without time zone,
		clientVersionSubmissionDate timestamp with time zone,
		serverVersionSubmissionDate timestamp with time zone,
		
		provider_id character varying(75),
		cityvillage character varying(75),
address1 character varying(75),
harikekf character varying(75),
tandavitaltdsistolik character varying(75),
highriskpostpartummaternalsepsis character varying(75),
highriskpostpartummastitis character varying(75),
integrasiprogramantitb character varying(75),
ancid character varying(75),
pncdate character varying(75),
pelayananfe character varying(75),
vitamina24jampp character varying(75),
highriskpostpartumpih character varying(75),
komplikasilainnya character varying(75),
treatment character varying(75),
rujukan character varying(75),
dirujukke character varying(75),
highriskpostpartuminfection character varying(75),
highriskpostpartumhemorrhage character varying(75),
integrasiprogramfotothorax character varying(75),
lokasiperiksa character varying(75),
tandavitaltddiastolik character varying(75),
tandavitalsuhu character varying(75),
vitamina2jampp character varying(75),
komplikasi character varying(75),
integrasiprogramantimalaria character varying(75));



