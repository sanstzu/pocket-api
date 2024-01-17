# Pocket API

REST API that judges code against a set of test cases using the [`pocket-watch`](#references) service.

> Currently only the code execution is supported 🥵

## Usage

<table>
    <tr>
        <th colspan="2">Route</th>
        <th>Method</th>
        <th>Description</th>
    </tr>
    <tr>
        <td>
            <code>/api</code>
        </td>
        <td>
            <code>/submit</code>
        </td>
        <td>
            <code>POST</code>
        </td>
        <td>
            Submit a source-code (<code>file</code>) to pocket using <code>form-data</code>
        </td>
    </tr>
</table>

## References

- [pocket-watch](https://github.com/joshjms/pocket-watch)
