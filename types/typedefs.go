package main

type FetchJSON struct {
	SampleCount int `json:"sample-count"`
	Results     []Result
	Disclaimer  string
}

type Result struct {
	snp_indel_silent_np     []Snp                   `json:"snp-indel-silent-np"`
	cnv_variants            []CnvVariants           `json:"cnv-variants"`
	snp_indel_silent        []Snp                   `json:"snp-indel-silent"`
	meta_data               MetaData                `json:"meta-data"`
	snp_indel_exonic        []Snp                   `json:"snp-indel-exonic"`
	cnv_intragenic_variants []CnvIntragenicVariants `json:"cnv-intragenic-variants"`
	sv_variants             []SvVariants            `json:"sv-variants"`
	snp_indel_exonic_np     []Snp                   `json:"snp-indel-exonic-np"`
}

type Snp struct {
	alt_allele            string
	oncokb_reported       int
	so_status_cv_id       int
	rlevel                string
	mafreq_1000g          string
	mrev_status_name      string
	variant_class         string
	confidence_cv_id      int
	so_status_name        string
	oncokb_ver            string
	mrev_status_cv_id     int
	ref_allele            string
	chromosome            string
	variant_class_cv_id   int
	snp_indel_tool_name   string
	cDNA_change           string
	occurance_in_pop      string
	so_comments           string
	normal_dp             int
	variant_status_name   string
	comments              string
	transcript_id         string
	variant_status_cv_id  int
	dmp_variant_id        int
	is_hotspot            int
	dbSNP_id              string
	treatments            string
	gene_id               string
	tumor_ad              int
	dmp_sample_mrev_id    int
	start_position        int
	snp_indel_variant_id  int
	occurance_in_normal   string
	mrev_comments         string
	normal_vfreq          float64
	level                 string
	tumor_vfreq           float64
	oncogenic             string
	exon_num              string
	normal_ad             int
	cosmic_id             string
	clinical_signed_out   string `json:"clinical-signed-out"`
	confidence_class      string
	dmp_sample_so_id      int
	tumor_dp              int
	oncokb_interpretation string
	aa_change             string
	is_reported           int
	d_tumor_ad            int
	d_tumor_rd            int
	d_tumor_dp            int
	d_tumor_vfreq         int
	s_tumor_ad            int
	s_tumor_dp            int
	s_tumor_rd            int
	s_tumor_vfreq         int
}

type CnvVariants struct {
	oncokb_reported       int
	gene_p_value          float64
	rlevel                int
	cnv_filter_cv_id      int
	confidence_cv_id      int
	oncokb_ver            string
	oncogenic             string
	chromosome            string
	variant_status_name   string
	comments              string
	variant_status_cv_id  int
	confidence_class      string
	treatments            string
	gene_fold_change      float64
	gene_id               string
	is_significant        int
	cytoband              string
	level                 string
	cnv_filter_name       string
	cnv_variant_id        int
	cnv_class_name        string
	clinical_signed_out   string `json:"clinical-signed-out"`
	cnv_class_cv_id       int
	oncokb_interpretation string
	is_reported           int
}

type MetaData struct {
	tmb_tt_percentile     float64
	tmb_score             float64
	dt_alys_end_time      string
	metastasis_site       string
	tmb_cohort_percentile float64
	mrev_status_name      string
	msi_score             string `json:"msi-score"`
	slide_viewer_id       string `json:"slide-viewer-id"`
	gene_panel            string `json:"gene-panel"`
	dmp_patient_id        string
	tmb_tt_cohort         float64
	so_comments           string
	dmp_sample_id         string
	retrieve_status       int
	cbx_sample_id         int
	dt_dms_start_time     string
	is_metastasis         int
	date_tumor_sequencing string
	somatic_status        string
	dmp_sample_so_id      int
	alys2sample_id        int
	tumor_type_code       string
	legacy_patient_id     string
	cbx_patient_id        int
	consent_parta         bool `json:"consent-parta"`
	consent_partc         bool `json:"consent-partc"`
	dmp_alys_task_id      int
	tumor_purity          string
	mrev_comments         string
	sample_coverage       int
	gender                int
	msi_type              string `json:"msi-type"`
	tumor_type_name       string
	legacy_sample_id      string
	outside_institute     string
	tmb_cohort            float64
	msi_comment           string `json:"msi-comment"`
	primary_site          string
	dmp_alys_task_name    string
	so_status_name        string
}

type CnvIntragenicVariants struct {
	refseq_acc            string
	treatments            int
	oncokb_reported       int
	cytoband              string
	level                 int
	rlevel                int
	cluster_2             string
	cnv_variant_id        int
	gene_id               string
	cluster_1             string
	confidence_cv_id      int
	oncokb_ver            int
	oncogenic             int
	oncokb_interpretation int
	variant_status_cv_id  int
}

type SvVariants struct {
	site1_desc              string
	oncokb_reported         int
	tumor_variant_count     int
	rlevel                  int
	normal_variant_count    int
	variant_status_name     string
	site1_gene              string
	site2_chrom             string
	oncogenic               string
	connection_type         string
	site2_desc              string
	event_info              string
	comments                string
	conn_type               string
	sv_variant_id           int
	mapq                    int
	site1_chrom             string
	site1_pos               int
	sv_desc                 string
	sv_class_name           string
	breakpoint_type         string
	site2_pos               int
	site2_gene              string
	treatments              string
	paired_end_read_support int
	tumor_read_count        int
	annotation              string
	split_read_support      int
	level                   string
	normal_read_count       int
	confidence_class        string
	oncokb_interpretation   string
	sv_length               int
}
