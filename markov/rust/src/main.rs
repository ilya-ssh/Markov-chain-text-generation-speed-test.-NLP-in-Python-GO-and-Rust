use rand::prelude::{IteratorRandom, SliceRandom};
use std::io;

const CUSTOM_PUNCTUATION: &[char] = &['.', '!', '?'];

fn build_markov_model(text: &str) -> std::collections::HashMap<&str, Vec<&str>> {
    let words: Vec<&str> = text.split_whitespace().collect();
    let mut model = std::collections::HashMap::new();

    for i in 0..words.len() - 1 {
        let current_word = words[i];
        let next_word = words[i + 1];

        model
            .entry(current_word)
            .or_insert_with(Vec::new)
            .push(next_word);
    }

    model
}

fn generate_text(model: &std::collections::HashMap<&str, Vec<&str>>, num_sentences: usize) -> String {
    let mut rng = rand::thread_rng();
    let mut generated_text = Vec::new();
    let mut current_word = if let Some(start_word_candidates) = model.keys().cloned().filter(|&word| word.chars().next().unwrap().is_uppercase()).collect::<Vec<_>>().choose(&mut rng) {
        start_word_candidates
    } else {
        model.keys().cloned().choose(&mut rng).unwrap()
    };

    let mut sentence_count = 0;

    while sentence_count < num_sentences {
        if let Some(next_word_options) = model.get(current_word) {
            let next_word = next_word_options.choose(&mut rng).unwrap();
            generated_text.push(*next_word);
            current_word = next_word;

            // Check for the presence of a punctuation mark
            if next_word.chars().last().map_or(false, |c| CUSTOM_PUNCTUATION.contains(&c)) {
                sentence_count += 1;
            }
        } else {
            break;
        }
    }

    generated_text.join(" ")
}

fn load_text_files(directory_path: &str, num_files: usize) -> Vec<String> {
    let mut texts = Vec::new();

    for i in 1..=num_files {
        // Format the file name with leading zeros
        let file_name = format!("{:03}.txt", i);
        let file_path = format!("{}/{}", directory_path, file_name);

        if let Ok(content) = std::fs::read_to_string(&file_path) {
            texts.push(content);
        } else {
            eprintln!("File not found: {}", file_path);
        }
    }

    texts
}

fn main() {
	
    let directory_path = "../../../corpus";
	let start_time = std::time::Instant::now();	
    let input_texts = load_text_files(directory_path, 511);
    let combined_text = input_texts.join(" ");
    let markov_model = build_markov_model(&combined_text);
    let generated_text = generate_text(&markov_model, 500000);
    let end_time = std::time::Instant::now();
    println!("Generated Text:");
    println!("{}", generated_text);
    let execution_time = end_time.duration_since(start_time);
    println!("Execution Time: {} seconds", execution_time.as_secs_f64());
	println!("Press Enter to exit...");
    let _ = io::stdin().read_line(&mut String::new());

}
