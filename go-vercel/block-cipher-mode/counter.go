package blockCipherMode

import (
	"ICCBES/lib"
	"ICCBES/lib/encryption_algorithm"
	"ICCBES/lib/utils"
)

func EncryptCTR(plainText []byte, key []byte, encryptionAlgorithm encryption_algorithm.EncryptionAlgorithm, iv []byte) []byte {
	// Initialize counter (can be a byte slice or integer)
	counter := iv

	// Split plainText into blocks (same size as key)
	plainTextBlocks := utils.TextToBlocks(plainText, len(key))
	blockLength := len(plainTextBlocks)
	cipherTextBlocks := make([][]byte, blockLength)

	for i := 0; i < blockLength; i++ {
		currentBlock := plainTextBlocks[i]

		// Generate keystream (encrypt counter)
		keystream := encryptCTRBlock(counter, key)

		// Encrypt block using XOR with keystream
		currentBlock = utils.DoBitXOR(currentBlock, keystream)

		// Save the result
		cipherTextBlocks[i] = currentBlock

		// Increment counter for next iteration (specific logic needed)
		counter = incrementCounter(counter) // Function to handle counter increment
	}

	// Merge blocks into one
	cipherText := make([]byte, len(plainText))
	for i := 0; i < blockLength; i++ {
		for j := 0; j < len(key); j++ {
			cipherText[i*len(key)+j] = cipherTextBlocks[i][j]
		}
	}
	return cipherText
}

func DecryptCTR(cipherText []byte, key []byte, decryptionAlgorithm lib.DecryptionAlgorithm, iv []byte) []byte {
	// Split cipherText into blocks (same size as key)
	cipherTextBlocks := utils.TextToBlocks(cipherText, len(key))
	blockLength := len(cipherTextBlocks)
	plainTextBlocks := make([][]byte, blockLength)
	counter := iv

	for i := 0; i < blockLength; i++ {
		currentBlock := cipherTextBlocks[i]

		// Generate keystream (encrypt counter)
		keystream := encryptCTRBlock(counter, encryptionAlgorithm)

		// Decrypt block using XOR with keystream
		currentBlock = utils.DoBitXOR(currentBlock, keystream)

		// Save the result
		plainTextBlocks[i] = currentBlock

		// Increment counter for next iteration (specific logic needed)
		counter = incrementCounter(counter) // Function to handle counter increment
	}

	// Merge blocks into one
	plainText := make([]byte, len(cipherText))
	for i := 0; i < blockLength; i++ {
		for j := 0; j < len(key); j++ {
			plainText[i*len(key)+j] = plainTextBlocks[i][j]
		}
	}
	return plainText
}

func encryptCTRBlock(counter []byte, encryptionAlgorithm encryption_algorithm.EncryptionAlgorithm) []byte {
	// Encrypt the counter block
	encryptedBlock := encryptionAlgorithm(counter, []byte{})

	// Return the entire encrypted block as keystream
	return encryptedBlock
}

// Implement incrementCounter function based on your counter type (byte slice or integer)
func incrementCounter(counter []byte) []byte {
	// Implement logic to increment the counter byte slice (e.g., carry over)
	// ...
	return counter
}