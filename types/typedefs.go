package types

type FetchJSON struct {
	SampleCount int `json:"sample-count"`
	Results     []Result
	Disclaimer  string
}

type Result struct {
	Snp_indel_silent_np     []Snp                   `json:"snp-indel-silent-np"`
	Cnv_variants            []CnvVariants           `json:"cnv-variants"`
	Snp_indel_silent        []Snp                   `json:"snp-indel-silent"`
	Meta_data               MetaData                `json:"meta-data"`
	Snp_indel_exonic        []Snp                   `json:"snp-indel-exonic"`
	Cnv_intragenic_variants []CnvIntragenicVariants `json:"cnv-intragenic-variants"`
	Sv_variants             []SvVariants            `json:"sv-variants"`
	Snp_indel_exonic_np     []Snp                   `json:"snp-indel-exonic-np"`
	Last_modified           any
}

type Snp struct {
	Alt_allele            string
	Oncokb_reported       int
	So_status_cv_id       int
	Rlevel                string
	Mafreq_1000g          string
	Mrev_status_name      string
	Variant_class         string
	Confidence_cv_id      int
	So_status_name        string
	Oncokb_ver            string
	Mrev_status_cv_id     int
	Ref_allele            string
	Chromosome            string
	Variant_class_cv_id   int
	Snp_indel_tool_name   string
	CDNA_change           string
	Occurance_in_pop      string
	So_comments           string
	Normal_dp             int
	Variant_status_name   string
	Comments              string
	Transcript_id         string
	Variant_status_cv_id  int
	Dmp_variant_id        int
	Is_hotspot            int
	DbSNP_id              string
	Treatments            string
	Gene_id               string
	Tumor_ad              int
	Dmp_sample_mrev_id    int
	Start_position        int
	Snp_indel_variant_id  int
	Occurance_in_normal   string
	Mrev_comments         string
	Normal_vfreq          float64
	Level                 string
	Tumor_vfreq           float64
	Oncogenic             string
	Exon_num              string
	Normal_ad             int
	Cosmic_id             string
	Clinical_signed_out   string `json:"clinical-signed-out"`
	Confidence_class      string
	Dmp_sample_so_id      int
	Tumor_dp              int
	Oncokb_interpretation string
	Aa_change             string
	Is_reported           int
	D_tumor_ad            int
	D_tumor_rd            int
	D_tumor_dp            int
	D_tumor_vfreq         int
	S_tumor_ad            int
	S_tumor_dp            int
	S_tumor_rd            int
	S_tumor_vfreq         int
}

type CnvVariants struct {
	Oncokb_reported       int
	Gene_p_value          float64
	Rlevel                int
	Cnv_filter_cv_id      int
	Confidence_cv_id      int
	Oncokb_ver            string
	Oncogenic             string
	Chromosome            string
	Variant_status_name   string
	Comments              string
	Variant_status_cv_id  int
	Confidence_class      string
	Treatments            string
	Gene_fold_change      float64
	Gene_id               string
	Is_significant        int
	Cytoband              string
	Level                 string
	Cnv_filter_name       string
	Cnv_variant_id        int
	Cnv_class_name        string
	Clinical_signed_out   string `json:"clinical-signed-out"`
	Cnv_class_cv_id       int
	Oncokb_interpretation string
	Is_reported           int
}

type MetaData struct {
	Tmb_tt_percentile     float64
	Tmb_score             float64
	Dt_alys_end_time      string
	Metastasis_site       string
	Tmb_cohort_percentile float64
	Mrev_status_name      string
	Msi_score             string `json:"msi-score"`
	Slide_viewer_id       string `json:"slide-viewer-id"`
	Gene_panel            string `json:"gene-panel"`
	Dmp_patient_id        string
	Tmb_tt_cohort         float64
	So_comments           string
	Dmp_sample_id         string
	Retrieve_status       int
	Cbx_sample_id         int
	Dt_dms_start_time     string
	Is_metastasis         int
	Date_tumor_sequencing string
	Somatic_status        string
	Dmp_sample_so_id      int
	Alys2sample_id        int
	Tumor_type_code       string
	Legacy_patient_id     string
	Cbx_patient_id        int
	Consent_parta         bool `json:"consent-parta"`
	Consent_partc         bool `json:"consent-partc"`
	Dmp_alys_task_id      int
	Tumor_purity          string
	Mrev_comments         string
	Sample_coverage       int
	Gender                int
	Msi_type              string `json:"msi-type"`
	Tumor_type_name       string
	Legacy_sample_id      string
	Outside_institute     string
	Tmb_cohort            float64
	Msi_comment           string `json:"msi-comment"`
	Primary_site          string
	Dmp_alys_task_name    string
	So_status_name        string
}

type CnvIntragenicVariants struct {
	Refseq_acc            string
	Treatments            int
	Oncokb_reported       int
	Cytoband              string
	Level                 int
	Rlevel                int
	Cluster_2             string
	Cnv_variant_id        int
	Gene_id               string
	Cluster_1             string
	Confidence_cv_id      int
	Oncokb_ver            int
	Oncogenic             int
	Oncokb_interpretation int
	Variant_status_cv_id  int
}

type SvVariants struct {
	Site1_desc              string
	Oncokb_reported         int
	Tumor_variant_count     int
	Rlevel                  int
	Normal_variant_count    int
	Variant_status_name     string
	Site1_gene              string
	Site2_chrom             string
	Oncogenic               string
	Connection_type         string
	Site2_desc              string
	Event_info              string
	Comments                string
	Conn_type               string
	Sv_variant_id           int
	Mapq                    int
	Site1_chrom             string
	Site1_pos               int
	Sv_desc                 string
	Sv_class_name           string
	Breakpoint_type         string
	Site2_pos               int
	Site2_gene              string
	Treatments              string
	Paired_end_read_support int
	Tumor_read_count        int
	Annotation              string
	Split_read_support      int
	Level                   string
	Normal_read_count       int
	Confidence_class        string
	Oncokb_interpretation   string
	Sv_length               int
}
