const integerOnly = (text) => {
    return text.replace(/[^0-9]/g, "");
};

const resort = (data, field, asc) => {
    return data.sort((a, b) => {
        const aValue = valueByPath(a, field);
        const bValue = valueByPath(b, field);
        if (aValue === bValue) {
            return 0;
        }

        if (asc === true) {
            return aValue > bValue ? 1 : -1;
        }

        return aValue > bValue ? -1 : 1;
    });
};

const valueByPath = (obj, path) => {
    return path
        .split(".")
        .reduce(
            (acc, key) =>
                acc && acc[key] !== "undefined" ? `${acc[key]}` : undefined,
            obj
        );
};

const handleOpenS3Object = (bucketName, fileName) => {
    const form = document.createElement("form");
    form.method = "POST";
    form.action = `${process.env.REACT_APP_API_URL}/api/s3/load`;
    form.target = "_blank";

    const dataInput1 = document.createElement("input");
    dataInput1.type = "hidden";
    dataInput1.name = "bucketName";
    dataInput1.value = bucketName;
    form.appendChild(dataInput1);

    const dataInput2 = document.createElement("input");
    dataInput2.type = "hidden";
    dataInput2.name = "fileName";
    dataInput2.value = fileName;
    form.appendChild(dataInput2);

    document.body.appendChild(form);
    form.submit();

    document.body.removeChild(form);
};

export { valueByPath, integerOnly, resort, handleOpenS3Object };
