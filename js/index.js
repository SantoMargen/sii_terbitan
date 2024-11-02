const CryptoJS = require('crypto-js');


const iv = CryptoJS.lib.WordArray.random(16);
const key = CryptoJS.enc.Utf8.parse('1n1S4ng4tr4has14');

// const plaintext = {"request":{"ID":76}}
const plaintext = {"request":
  {
    "ID":59,
    "title": "Sample Paper Update",
    "authors": "Author One Update",
    "co_authors": "Coauthor A, Coauthor B Update",
    "publication_date": "2023-11-01T12:34:56Z",
    "journal": "Journal of Testing Update",
    "volume": 11,
    "issue": 4,
    "page_range": "100-110 Update",
    "doi": "10.1000/sample.doi Update",
    "abstract": "This is a sample abstract. Update",
    "keywords": "sample, paper, test Update",
    "research_type": "Empirical Update",
    "funding_info": "Some Funding Update",
    "affiliations": "Some UniversityUpdate ",
    "full_text_link": "http://example.com/fulltext Update",
    "language": "English Update",
    "license": "CC BY Update",
    "notes": "Additional notes here.Update",
  }
}

// }
// const plaintext = {"request":
//     {"category_name": "Ilussion", "description": "123 Main St, City A, Country A", "entry_user": "Hisbikal","ID":22}
//   }


function Encrypt(plaintext) {
  const strText = JSON.stringify(plaintext);

  const encrypted = CryptoJS.AES.encrypt(strText, key, {
    iv: iv,
    mode: CryptoJS.mode.CBC,
    padding: CryptoJS.pad.Pkcs7
  });

  const ciphertext = iv.concat(encrypted.ciphertext).toString(CryptoJS.enc.Base64);
  return ciphertext
}

function Decrypt(ciphertext) {
  const ciphertextDecoded = CryptoJS.enc.Base64.parse(ciphertext);
  const iv = ciphertextDecoded.clone();
  iv.sigBytes = 16;
  iv.clamp();

  const encryptedMessage = ciphertextDecoded.clone();
  encryptedMessage.words.splice(0, 4);
  encryptedMessage.sigBytes -= 16;

  const decrypted = CryptoJS.AES.decrypt(
    {
      ciphertext: encryptedMessage,
    },
    key,
    {
      iv: iv,
      mode: CryptoJS.mode.CBC,
      padding: CryptoJS.pad.Pkcs7,
    }
  );

  const plaintext = decrypted.toString(CryptoJS.enc.Utf8);
  return plaintext;
}



console.log("START =======================");

console.log("ENCRYPT: ", Encrypt(plaintext));
// console.log("DECRYPT: ", Decrypt("StHMXv9196HvBzkHGI+sI0CgTC5D3TrE//SA12bGWiO2jsxHpmpffIIP/WIxZYua"))