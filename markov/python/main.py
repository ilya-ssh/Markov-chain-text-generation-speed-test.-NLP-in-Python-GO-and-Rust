import random
import time


custom_punctuation = {'.', '!', '?'}

def build_markov_model(text):
    words = text.split()
    model = {}

    for i in range(len(words) - 1):
        current_word = words[i]
        next_word = words[i + 1]

        if current_word not in model:
            model[current_word] = []

        model[current_word].append(next_word)

    return model

def generate_text(model, num_sentences=1, start_word=None):
    if start_word is None:
        # Select a capitalized word as the starting point
        start_word_candidates = [word for word in model.keys() if word[0].isupper()]
        start_word = random.choice(start_word_candidates) if start_word_candidates else random.choice(list(model.keys()))

    current_word = start_word
    generated_text = [current_word]

    sentence_count = 0

    while sentence_count < num_sentences:
        if current_word in model:
            next_word = random.choice(model[current_word])
            generated_text.append(next_word)
            current_word = next_word

            # Check for the presence of a punctuation mark
            if next_word[-1] in custom_punctuation:
                sentence_count += 1
        else:
            break

    return ' '.join(generated_text)

def load_text_files(directory_path, num_files=511):
    texts = []

    for i in range(1, num_files + 1):
        # Format the file name with leading zeros
        file_name = f"{i:03d}.txt"
        file_path = f"{directory_path}/{file_name}"

        try:
            with open(file_path, 'r', encoding='utf-8') as file:
                text = file.read()
                texts.append(text)
        except FileNotFoundError:
            print(f"File not found: {file_path}")

    return texts

if __name__ == "__main__":
    directory_path = "../corpus"
    start_time = time.time()
    input_texts = load_text_files(directory_path)

    # Combine all loaded texts into a single string
    combined_text = ' '.join(input_texts)

    # Build the Markov model
    markov_model = build_markov_model(combined_text)

    # Measure execution time for generating text
    

    # Generate text using the model with 3 sentences as an example
    generated_text = generate_text(markov_model, num_sentences=500000)

    end_time = time.time()

    print("Generated Text:")
    print(generated_text)

    # Calculate and print the execution time
    execution_time = end_time - start_time
    print(f"Execution Time: {execution_time} seconds")
    input()
