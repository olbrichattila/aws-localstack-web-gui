
const Layout = ({ Component, Icon, title }) => {

    return (
        <>
            <div className="containerHeader">
                <Icon /><span>{title}</span>
            </div>
            <div className="pageContainer">
                <Component />
            </div>
        </>
    );

}

export default Layout;
