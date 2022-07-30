package genome

type Genome struct {
	Variant         string `json:"variant"`
	GenomicLocation struct {
		Chromosome      string `json:"chromosome"`
		Start           int    `json:"start"`
		End             int    `json:"end"`
		ReferenceAllele string `json:"referenceAllele"`
		VariantAllele   string `json:"variantAllele"`
	} `json:"genomicLocation"`
	StrandSign             string `json:"strandSign"`
	VariantType            string `json:"variantType"`
	AssemblyName           string `json:"assemblyName"`
	CanonicalTranscriptID  string `json:"canonicalTranscriptId"`
	TranscriptConsequences []struct {
		TranscriptID     string `json:"transcriptId"`
		CodonChange      string `json:"codonChange,omitempty"`
		AminoAcids       string `json:"aminoAcids,omitempty"`
		AminoAcidRef     string `json:"aminoAcidRef,omitempty"`
		AminoAcidAlt     string `json:"aminoAcidAlt,omitempty"`
		EntrezGeneID     string `json:"entrezGeneId"`
		ConsequenceTerms string `json:"consequenceTerms"`
		HugoGeneSymbol   string `json:"hugoGeneSymbol"`
		HgvspShort       string `json:"hgvspShort"`
		Hgvsp            string `json:"hgvsp,omitempty"`
		Hgvsc            string `json:"hgvsc"`
		ProteinPosition  struct {
			Start int `json:"start"`
			End   int `json:"end"`
		} `json:"proteinPosition,omitempty"`
		RefSeq                string  `json:"refSeq,omitempty"`
		VariantClassification string  `json:"variantClassification"`
		Exon                  string  `json:"exon,omitempty"`
		PolyphenScore         float64 `json:"polyphenScore,omitempty"`
		PolyphenPrediction    string  `json:"polyphenPrediction,omitempty"`
		SiftScore             float64 `json:"siftScore,omitempty"`
		SiftPrediction        string  `json:"siftPrediction,omitempty"`
		UniprotID             string  `json:"uniprotId,omitempty"`
	} `json:"transcriptConsequences"`
	TranscriptConsequenceSummaries []struct {
		TranscriptID     string `json:"transcriptId"`
		CodonChange      string `json:"codonChange,omitempty"`
		AminoAcids       string `json:"aminoAcids,omitempty"`
		AminoAcidRef     string `json:"aminoAcidRef,omitempty"`
		AminoAcidAlt     string `json:"aminoAcidAlt,omitempty"`
		EntrezGeneID     string `json:"entrezGeneId"`
		ConsequenceTerms string `json:"consequenceTerms"`
		HugoGeneSymbol   string `json:"hugoGeneSymbol"`
		HgvspShort       string `json:"hgvspShort"`
		Hgvsp            string `json:"hgvsp,omitempty"`
		Hgvsc            string `json:"hgvsc"`
		ProteinPosition  struct {
			Start int `json:"start"`
			End   int `json:"end"`
		} `json:"proteinPosition,omitempty"`
		RefSeq                string  `json:"refSeq,omitempty"`
		VariantClassification string  `json:"variantClassification"`
		Exon                  string  `json:"exon,omitempty"`
		PolyphenScore         float64 `json:"polyphenScore,omitempty"`
		PolyphenPrediction    string  `json:"polyphenPrediction,omitempty"`
		SiftScore             float64 `json:"siftScore,omitempty"`
		SiftPrediction        string  `json:"siftPrediction,omitempty"`
		UniprotID             string  `json:"uniprotId,omitempty"`
	} `json:"transcriptConsequenceSummaries"`
	TranscriptConsequenceSummary struct {
		TranscriptID     string `json:"transcriptId"`
		CodonChange      string `json:"codonChange"`
		AminoAcids       string `json:"aminoAcids"`
		AminoAcidRef     string `json:"aminoAcidRef"`
		AminoAcidAlt     string `json:"aminoAcidAlt"`
		EntrezGeneID     string `json:"entrezGeneId"`
		ConsequenceTerms string `json:"consequenceTerms"`
		HugoGeneSymbol   string `json:"hugoGeneSymbol"`
		HgvspShort       string `json:"hgvspShort"`
		Hgvsp            string `json:"hgvsp"`
		Hgvsc            string `json:"hgvsc"`
		ProteinPosition  struct {
			Start int `json:"start"`
			End   int `json:"end"`
		} `json:"proteinPosition"`
		RefSeq                string  `json:"refSeq"`
		VariantClassification string  `json:"variantClassification"`
		Exon                  string  `json:"exon"`
		PolyphenScore         float64 `json:"polyphenScore"`
		PolyphenPrediction    string  `json:"polyphenPrediction"`
		SiftScore             float64 `json:"siftScore"`
		SiftPrediction        string  `json:"siftPrediction"`
		UniprotID             string  `json:"uniprotId"`
	} `json:"transcriptConsequenceSummary,omitempty"`
	TranscriptConsequenceSummary0 struct {
		TranscriptID     string `json:"transcriptId"`
		CodonChange      string `json:"codonChange"`
		AminoAcids       string `json:"aminoAcids"`
		AminoAcidRef     string `json:"aminoAcidRef"`
		AminoAcidAlt     string `json:"aminoAcidAlt"`
		EntrezGeneID     string `json:"entrezGeneId"`
		ConsequenceTerms string `json:"consequenceTerms"`
		HugoGeneSymbol   string `json:"hugoGeneSymbol"`
		HgvspShort       string `json:"hgvspShort"`
		Hgvsp            string `json:"hgvsp"`
		Hgvsc            string `json:"hgvsc"`
		ProteinPosition  struct {
			Start int `json:"start"`
			End   int `json:"end"`
		} `json:"proteinPosition"`
		RefSeq                string `json:"refSeq"`
		VariantClassification string `json:"variantClassification"`
		Exon                  string `json:"exon"`
		UniprotID             string `json:"uniprotId"`
	} `json:"transcriptConsequenceSummary,omitempty"`
}
