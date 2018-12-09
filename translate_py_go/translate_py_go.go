package main

import (
"fmt"
"io/ioutil"
"strings"
//"reflect"
)


func contains(arr []string, str string) bool {
   for _, a := range arr {
      if a == str {
         return true
      }
   }
   return false
}

func create_lookup_tables(text []string) map[string]int{

	var vocab_to_int map[string]int
	vocab_to_int = make(map[string]int)

	var list []string
	for i := 0; i < len(text); i++ {
		words := strings.Split(text[i], " ")
		for j := 0; j < len(words); j++ {
			if contains(list,words[j]) == false {
			list = append(list,words[j])
			}
		}
	}
    
    //fmt.Println(list)
    //fmt.Println(len(list))

    //'<PAD>': 0, '<EOS>': 1, '<UNK>': 2, '<GO>': 3

    vocab_to_int["<PAD>"] = 0
    vocab_to_int["<EOS>"] = 1
    vocab_to_int["<UNK>"] = 2
    vocab_to_int["<GO>"] = 3

    for i := 0; i < len(list); i++ {
		vocab_to_int[list[i]] = i+4	
	}
    return vocab_to_int
}


func reverseMap(m map[string]int) map[int]string {
    n := make(map[int]string)
    for k, v := range m {
        n[v] = k
    }
    return n
}


func en_text_to_ids(source_text []string,  source_vocab_to_int map[string]int) [][]int{
	var sentences [][]int
	var sentence []int
	for i := 0; i < len(source_text); i++ {
		sentence = sentence[:0]
		words := strings.Split(source_text[i], " ")
		for j := 0; j < len(words); j++ {
			sentence = append(sentence,source_vocab_to_int[words[j]])
		}
		sentence = append(sentence,source_vocab_to_int["<EOS>"])
		sentences = append(sentences,sentence)
	}
	return sentences
}

func fr_text_to_ids(target_text []string, target_vocab_to_int map[string]int) [][]int{
	var sentences [][]int
	var sentence []int
	for i := 0; i < len(target_text); i++ {
		sentence = sentence[:0]
		words := strings.Split(target_text[i], " ")
		for j := 0; j < len(words); j++ {
			sentence = append(sentence,target_vocab_to_int[words[j]])
		}
		sentence = append(sentence,target_vocab_to_int["<EOS>"])
		sentences = append(sentences,sentence)
	}
	return sentences

}


func main() {
	//fmt.Println("Hello, world.")

	
	
	var string_data string

	data, err := ioutil.ReadFile("small_vocab_en")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    //fmt.Println("Contents of file:", string(data))
    string_data = string(data)
    english_sentences := strings.Split(string_data, "\n")
    

	data, err = ioutil.ReadFile("small_vocab_fr")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    string_data = string(data)
    french_sentences := strings.Split(string_data, "\n")
    //fmt.Println(reflect.TypeOf(english_sentences))
    //fmt.Println(len(english_sentences))

    //fmt.Println(english_sentences[0])

    var english_id_table map[string]int
    var id_english_table map[int]string

	var french_id_table map[string]int
    var id_french_table map[int]string

    english_id_table = create_lookup_tables(english_sentences)
    id_english_table = reverseMap(english_id_table)

    french_id_table = create_lookup_tables(french_sentences)
    id_french_table = reverseMap(french_id_table)

    _ = id_french_table
    _ = id_english_table
    //fmt.Println(english_id_table,"\n",french_id_table)
    //fmt.Println(id_english_table,"\n",id_french_table)

    var english_sentences_ids [][]int
    english_sentences_ids = en_text_to_ids(english_sentences,english_id_table)
    fmt.Println(english_sentences_ids)


    var french_sentences_ids [][]int
    french_sentences_ids = fr_text_to_ids(french_sentences,french_id_table)
    fmt.Println(french_sentences_ids)

    //the united states is usually chilly during july , and it is usually freezing in november .
    //fmt.Println(english_id_table["the"])
    //fmt.Println(english_id_table["united"])
    //fmt.Println(english_id_table["states"])
    //fmt.Println(english_id_table["is"])
    //fmt.Println(english_id_table["usually"])
    //fmt.Println(english_id_table["chilly"])
    //fmt.Println(english_id_table["during"])
    //fmt.Println(english_id_table["july"])
	//fmt.Println(english_id_table["and"])
	//fmt.Println(english_id_table["it"])
	//fmt.Println(english_id_table["is"])
	//fmt.Println(english_id_table["usually"])
	//fmt.Println(english_id_table["freezing"])
	//fmt.Println(english_id_table["in"])    
	//fmt.Println(english_id_table["november"])


}