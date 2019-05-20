//mempermudah pembuatan elemen
const cE = React.createElement

class FixedForm extends React.Component{
    //ctor
    constructor(props){
        super(props);
        this.state = {number: '', text: ''};

        this.handleChangeNum = this.handleChangeNum.bind(this);
        this.handleChangeTxt = this.handleChangeTxt.bind(this);
        this.IntToString = this.IntToString.bind(this);
        this.StringToInt = this.StringToInt.bind(this);
    }

    handleChangeNum(event){
        //pengubahan isi textarea mengubah isi dari this.state.number
        this.setState({number: event.target.value, text: this.state.text})
    }

    handleChangeTxt(event){
        //pengubahan isi textarea mengubah isi dari this.state.text
        this.setState({number: this.state.number, text: event.target.value});
    }

    IntToString(){
        //GET -> mengambil angka dari number textarea dan alert hasil convert
        var request = new XMLHttpRequest();
        request.open('GET', 'http://localhost:8080/spell?number=' + this.state.number, true);
        request.onload = function () {
            var data = JSON.parse(this.response)
            if (request.status >= 200 && request.status < 400){
                //alert hasil
                alert(data.text);
            } else {
                alert("error");
            }
        }
        request.send();
    }

    StringToInt(){
        //POST -> mengambil teks dari text textarea dan alert hasil convert
        var request = new XMLHttpRequest();
        request.open('POST', 'http://localhost:8080/read?text=' + this.state.text, true);
        request.onload = function (){
            var data = JSON.parse(this.response)
            if (request.status >= 200 && request.status < 400){
                //alert hasil
                alert(data.number);
            } else {
                alert("error")
            }
        }
        request.send();
    }

    render(){
        //UI nya susah weh kok gw cupu sih?
        return cE('div', null, [
            cE(
                'h2',
                { key: "1" },
                "Number"
            ),
            //textarea flexibel ukuran
            cE(
                'textarea',
                { key: "ta1", value: this.state.number,
                  onChange: this.handleChangeNum},
                'Number'
            ),
            cE(
                'h4',
                { key: "nl"},
                ''
            ),
            cE(
                'button',
                { key: "bt1", onClick: this.IntToString},
                'Convert to String'
            ),
            cE(
                'h3',
                { key: "n2"},
                '=============================='
            ),
            cE(
                'h2',
                { key: "n3"},
                "Text"
            ),
            cE(
                'textarea',
                { key: "ta2", value: this.state.text,
                  onChange: this.handleChangeTxt},
                'Text'
            ),
            cE(
                'h4',
                { key: "n4"},
                ''
            ),
            cE(
                'button',
                { key: "bt2", onClick: this.StringToInt},
                'Convert to Number'
            )]
        );
    }
}

ReactDOM.render(
    cE(FixedForm),
    document.querySelector('#root')//di div root pada main.html
);