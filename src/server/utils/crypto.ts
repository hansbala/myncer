import * as crypto from "crypto";

export const encryptUsingEnvKey = (plaintext: string): {
    encryptedText: string,
    iv: string
} => {
    const encryptionKey = process.env.ENCRYPTION_KEY
    if (!encryptionKey) {
        throw new Error("ENCRYPTION_KEY not set")
    }
    const iv = crypto.randomBytes(16)

    const cipher = crypto.createCipheriv('aes-256-cbc', Buffer.from(encryptionKey, 'base64'), iv);

    // Update the cipher with the text to encrypt
    let encryptedText = cipher.update(plaintext);
    // Finalize the encryption
    encryptedText = Buffer.concat([encryptedText, cipher.final()]);

    return {
        encryptedText: encryptedText.toString('hex'),
        iv: iv.toString('hex')
    }
}

export const decryptUsingEnvKey = (hexEncryptedText: string, hexIv: string): string => {
    const encryptionKey = process.env.ENCRYPTION_KEY
    if (!encryptionKey) {
        throw new Error("ENCRYPTION_KEY not set")
    }
    const iv = Buffer.from(hexIv, 'hex')
    const encryptedText = Buffer.from(hexEncryptedText, 'hex')

    const decipher = crypto.createDecipheriv('aes-256-cbc', Buffer.from(encryptionKey, 'base64'), iv);
    // Update the decipher with the encrypted text
    let decrypted = decipher.update(encryptedText);
    // Finalize the decryption
    decrypted = Buffer.concat([decrypted, decipher.final()]);

    return decrypted.toString();
}

export const getRandomString = (length: number): string => {
    const validCharacters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let array = new Uint8Array(length);
    crypto.getRandomValues(array);
    array = array.map(x => validCharacters.charCodeAt(x % validCharacters.length));
    return String.fromCharCode(...array);
}