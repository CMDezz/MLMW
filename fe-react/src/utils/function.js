export const createSampleFileFormUrl = (url, fileType) => {
  // Regular expression to extract the filename
  const regex = /[^/]+$/;

  // Extract the filename
  const filename = url.match(regex)[0];
  const fileContent = ''; // Empty content for this example

  // Create the file object
  const file = new File([fileContent], filename, { type: fileType });
  return file;
};

export const normFile = (e) => {
  if (Array.isArray(e)) {
    return e;
  }
  return e && e.fileList;
};
