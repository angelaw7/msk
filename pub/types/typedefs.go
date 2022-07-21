package types

type FetchJSON struct {
	SampleCount int `json:"sample-count"`
	Results     []Result
	Disclaimer  string `json:"disclaimer"`
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
	Last_modified           any                     `json:"-"`
}

type Snp struct {
	Alt_allele            string  `json:"alt_allele"`
	Oncokb_reported       int     `json:"oncokb_reported"`
	So_status_cv_id       int     `json:"so_status_cv_id"`
	Rlevel                string  `json:"rlevel"`
	Mafreq_1000g          string  `json:"mafreq_1000g"`
	Mrev_status_name      string  `json:"mrev_status_name"`
	Variant_class         string  `json:"variant_class"`
	Confidence_cv_id      int     `json:"confidence_cv_id"`
	So_status_name        string  `json:"so_status_name"`
	Oncokb_ver            string  `json:"oncokb_ver"`
	Mrev_status_cv_id     int     `json:"mrev_status_cv_id"`
	Ref_allele            string  `json:"ref_allele"`
	Chromosome            string  `json:"chromosome"`
	Variant_class_cv_id   int     `json:"variant_class_cv_id"`
	Snp_indel_tool_name   string  `json:"snp_indel_tool_name"`
	CDNA_change           string  `json:"cDNA_change"`
	Occurance_in_pop      string  `json:"occurance_in_pop"`
	So_comments           string  `json:"so_comments"`
	Normal_dp             int     `json:"normal_dp"`
	Variant_status_name   string  `json:"variant_status_name"`
	Comments              string  `json:"comments"`
	Transcript_id         string  `json:"transcript_id"`
	Variant_status_cv_id  int     `json:"variant_status_cv_id"`
	Dmp_variant_id        int     `json:"dmp_variant_id"`
	Is_hotspot            int     `json:"is_hotspot"`
	DbSNP_id              string  `json:"dbSNP_id"`
	Treatments            string  `json:"treatments"`
	Gene_id               string  `json:"gene_id"`
	Tumor_ad              int     `json:"tumor_ad"`
	Dmp_sample_mrev_id    int     `json:"dmp_sample_mrev_id"`
	Start_position        int     `json:"start_position"`
	Snp_indel_variant_id  int     `json:"snp_indel_variant_id"`
	Occurance_in_normal   string  `json:"occurance_in_normal"`
	Mrev_comments         string  `json:"mrev_comments"`
	Normal_vfreq          float64 `json:"normal_vfreq"`
	Level                 string  `json:"level"`
	Tumor_vfreq           float64 `json:"tumor_vfreq"`
	Oncogenic             string  `json:"oncogenic"`
	Exon_num              string  `json:"exon_num"`
	Normal_ad             int     `json:"normal_ad"`
	Cosmic_id             string  `json:"cosmic_id"`
	Clinical_signed_out   string  `json:"clinical-signed-out"`
	Confidence_class      string  `json:"confidence_class"`
	Dmp_sample_so_id      int     `json:"dmp_sample_so_id"`
	Tumor_dp              int     `json:"tumor_dp"`
	Oncokb_interpretation string  `json:"oncokb_interpretation"`
	Aa_change             string  `json:"aa_change"`
	Is_reported           int     `json:"is_reported"`
	D_tumor_ad            any     `json:"d_tumor_ad"`
	D_tumor_rd            any     `json:"d_tumor_rd"`
	D_tumor_dp            any     `json:"d_tumor_dp"`
	D_tumor_vfreq         any     `json:"d_tumor_vfreq"`
	S_tumor_ad            any     `json:"s_tumor_ad"`
	S_tumor_dp            any     `json:"s_tumor_dp"`
	S_tumor_rd            any     `json:"s_tumor_rd"`
	S_tumor_vfreq         any     `json:"s_tumor_vfreq"`
}

type CnvVariants struct {
	Oncokb_reported       int     `json:"oncokb_reported"`
	Gene_p_value          float64 `json:"gene_p_value"`
	Rlevel                any     `json:"rlevel"`
	Cnv_filter_cv_id      int     `json:"cnv_filter_cv_id"`
	Confidence_cv_id      int     `json:"confidence_cv_id"`
	Oncokb_ver            string  `json:"oncokb_ver"`
	Oncogenic             string  `json:"oncogenic"`
	Chromosome            string  `json:"chromosome"`
	Variant_status_name   string  `json:"variant_status_name"`
	Comments              string  `json:"comments"`
	Variant_status_cv_id  int     `json:"variant_status_cv_id"`
	Confidence_class      string  `json:"confidence_class"`
	Treatments            string  `json:"treatments"`
	Gene_fold_change      float64 `json:"gene_fold_change"`
	Gene_id               string  `json:"gene_id"`
	Is_significant        int     `json:"is_significant"`
	Cytoband              string  `json:"cytoband"`
	Level                 string  `json:"level"`
	Cnv_filter_name       string  `json:"cnv_filter_name"`
	Cnv_variant_id        int     `json:"cnv_variant_id"`
	Cnv_class_name        string  `json:"cnv_class_name"`
	Clinical_signed_out   string  `json:"clinical-signed-out"`
	Cnv_class_cv_id       int     `json:"cnv_class_cv_id"`
	Oncokb_interpretation string  `json:"oncokb_interpretation"`
	Is_reported           int     `json:"is_reported"`
}

type MetaData struct {
	Tmb_tt_percentile     float64 `json:"tmb_tt_percentile"`
	Tmb_score             float64 `json:"tmb_score"`
	Dt_alys_end_time      string  `json:"dt_alys_end_time"`
	Metastasis_site       string  `json:"metastasis_site"`
	Tmb_cohort_percentile float64 `json:"tmb_cohort_percentile"`
	Mrev_status_name      string  `json:"mrev_status_name"`
	Msi_score             string  `json:"msi-score"`
	Slide_viewer_id       string  `json:"slide-viewer-id"`
	Gene_panel            string  `json:"gene-panel"`
	Dmp_patient_id        string  `json:"dmp_patient_id"`
	Tmb_tt_cohort         float64 `json:"tmb_tt_cohort"`
	So_comments           string  `json:"so_comments"`
	Dmp_sample_id         string  `json:"dmp_sample_id"`
	Retrieve_status       int     `json:"retrieve_status"`
	Cbx_sample_id         int     `json:"cbx_sample_id"`
	Dt_dms_start_time     string  `json:"dt_dms_start_time"`
	Is_metastasis         int     `json:"is_metastasis"`
	Date_tumor_sequencing string  `json:"date_tumor_sequencing"`
	Somatic_status        string  `json:"somatic_status"`
	Dmp_sample_so_id      int     `json:"dmp_sample_so_id"`
	Alys2sample_id        int     `json:"alys2sample_id"`
	Tumor_type_code       string  `json:"tumor_type_code"`
	Legacy_patient_id     string  `json:"legacy_patient_id"`
	Cbx_patient_id        int     `json:"cbx_patient_id"`
	Consent_parta         bool    `json:"consent-parta"`
	Consent_partc         bool    `json:"consent-partc"`
	Dmp_alys_task_id      int     `json:"dmp_alys_task_id"`
	Tumor_purity          string  `json:"tumor_purity"`
	Mrev_comments         string  `json:"mrev_comments"`
	Sample_coverage       int     `json:"sample_coverage"`
	Gender                int     `json:"gender"`
	Msi_type              string  `json:"msi-type"`
	Tumor_type_name       string  `json:"tumor_type_name"`
	Legacy_sample_id      string  `json:"legacy_sample_id"`
	Outside_institute     string  `json:"outside_institute"`
	Tmb_cohort            float64 `json:"tmb_cohort"`
	Msi_comment           string  `json:"msi-comment"`
	Primary_site          string  `json:"primary_site"`
	Dmp_alys_task_name    string  `json:"dmp_alys_task_name"`
	So_status_name        string  `json:"so_status_name"`
}

type CnvIntragenicVariants struct {
	Refseq_acc            string `json:"refseq_acc"`
	Treatments            any    `json:"treatments"`
	Oncokb_reported       int    `json:"oncokb_reported"`
	Cytoband              string `json:"cytoband"`
	Level                 any    `json:"level"`
	Rlevel                any    `json:"rlevel"`
	Cluster_2             string `json:"cluster_2"`
	Cnv_variant_id        int    `json:"cnv_variant_id"`
	Gene_id               string `json:"gene_id"`
	Cluster_1             string `json:"cluster_1"`
	Confidence_cv_id      int    `json:"confidence_cv_id"`
	Oncokb_ver            any    `json:"oncokb_ver"`
	Oncogenic             any    `json:"oncogenic"`
	Oncokb_interpretation any    `json:"oncokb_interpretation"`
	Variant_status_cv_id  int    `json:"variant_status_cv_id"`
}

type SvVariants struct {
	Site1_desc              string `json:"site1_desc"`
	Oncokb_reported         int    `json:"oncokb_reported"`
	Tumor_variant_count     int    `json:"tumor_variant_count"`
	Rlevel                  any    `json:"rlevel"`
	Normal_variant_count    int    `json:"normal_variant_count"`
	Variant_status_name     string `json:"variant_status_name"`
	Site1_gene              string `json:"site1_gene"`
	Site2_chrom             string `json:"site2_chrom"`
	Oncogenic               string `json:"oncogenic"`
	Connection_type         string `json:"connection_type"`
	Site2_desc              string `json:"site2_desc"`
	Event_info              string `json:"event_info"`
	Comments                string `json:"comments"`
	Conn_type               string `json:"conn_type"`
	Sv_variant_id           int    `json:"sv_variant_id"`
	Mapq                    int    `json:"mapq"`
	Site1_chrom             string `json:"site1_chrom"`
	Site1_pos               int    `json:"site1_pos"`
	Sv_desc                 string `json:"sv_desc"`
	Sv_class_name           string `json:"sv_class_name"`
	Breakpoint_type         string `json:"breakpoint_type"`
	Site2_pos               int    `json:"site2_pos"`
	Site2_gene              string `json:"site2_gene"`
	Treatments              string `json:"treatments"`
	Paired_end_read_support int    `json:"paired_end_read_support"`
	Tumor_read_count        int    `json:"tumor_read_count"`
	Annotation              string `json:"annotation"`
	Split_read_support      int    `json:"split_read_support"`
	Level                   string `json:"level"`
	Normal_read_count       int    `json:"normal_read_count"`
	Confidence_class        string `json:"confidence_class"`
	Oncokb_interpretation   string `json:"oncokb_interpretation"`
	Sv_length               int    `json:"sv_length"`
}
