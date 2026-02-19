Why we use Argon2id by default

While bcrypt has been the industry standard for decades, Argon2id (the winner of the Password Hashing Competition) is our default for three critical reasons:

    Memory Hardness (GPU Resistance): Unlike bcrypt, which only requires CPU power, Argon2id is designed to be "memory-hard." This makes it physically and financially expensive for attackers to use high-powered GPUs or specialized hacking hardware (ASICs) to crack your users' passwords.

    Hybrid Protection: The id variant of Argon2 is a hybrid that protects against both GPU-based attacks and "side-channel" timing attacks, providing the most balanced security profile available in 2026.

    No Character Limits: bcrypt has a hard limit of 72 characters—anything longer is silently ignored. Argon2id has no such limit, allowing your users to use long, secure passphrases without compromising entropy.

Note: We provide a Hasher interface if your project specifically requires Bcrypt for legacy compatibility.